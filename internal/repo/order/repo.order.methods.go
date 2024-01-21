package order

import (
	"context"

	v1 "github.com/huseinnashr/pforder-backend/api/v1"
	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/pagination"
	"github.com/huseinnashr/pforder-backend/internal/pkg/querybuilder"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
	"github.com/lib/pq"
)

func (r *Repo) ListOrder(ctx context.Context, params domain.ListOrderParam) ([]domain.Order, string, error) {
	ctx, span := tracer.Start(ctx, "repo.ListOrder")
	defer span.End()

	qb := querybuilder.New(`
		select 
			o.id,
			o.order_name,
			oi.products,
			cc.company_name, 
			c.name,
			o.created_at,
			coalesce(oi.total_amount, 0) as total_amount,
			coalesce(oi.delivered_amount, 0) as delieved_amount
		from orders o 
		left join (
			select 
				oii.order_id, 
				array_agg(oii.product) as products, 
				sum(oii.price_per_unit * oii.quantity) as total_amount,
				sum(oii.price_per_unit * d.delivered_quantity) as delivered_amount
			from order_items oii
			left join (
				select 
					order_item_id,
					sum(delivered_quantity) as delivered_quantity
				from deliveries
				group by order_item_id
			) d on d.order_item_id = oii.id 
			group by oii.order_id
		) oi on oi.order_id = o.id
		join customers as c on c.user_id = o.customer_id
		join customer_companies cc on cc.company_id = c.company_id
		where 1 = 1
	`)

	if params.StartDate.Unix() != 0 {
		qb.AddQuery("AND created_at >= ?", params.StartDate)
	}

	if params.EndDate.Unix() != 0 {
		qb.AddQuery("AND created_at <= ?", params.EndDate)
	}

	if params.Search != "" {
		qb.AddQuery("AND (o.order_name = ? OR ? = any(oi.products))", params.Search, params.Search)
	}

	if params.Cursor != "" {
		cursorData := ListOrderCursor{}
		if err := pagination.DecodeCursor(params.Cursor, &cursorData); err != nil {
			return nil, "", err
		}
		if params.OrderType == v1.OrderType_ORDERTYPE_DESC {
			qb.AddQuery("AND (created_at, id) <= (?, ?)", cursorData.CreatedAt, cursorData.ID)
		} else {
			qb.AddQuery("AND (created_at, id) >= (?, ?)", cursorData.CreatedAt, cursorData.ID)
		}
	}

	if params.OrderType == v1.OrderType_ORDERTYPE_DESC {
		qb.AddString("ORDER BY (created_at, id) DESC")
	} else {
		qb.AddString("ORDER BY (created_at, id) ASC")
	}

	qb.AddQuery("LIMIT ?", params.PageSize+1)

	rows, err := r.sqlDatabase.QueryContext(ctx, qb.GetQuery(), qb.GetArgs()...)
	if err != nil {
		return nil, "", err
	}

	orders := []domain.Order{}
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(
			&order.ID, &order.OrderName, (*pq.StringArray)(&order.Products), &order.CustomerCompanyName,
			&order.CustomerName, &order.OrderDate, &order.TotalAmount, &order.DeliveredAmount,
		)
		if err != nil {
			return nil, "", err
		}

		orders = append(orders, order)
	}

	cursor := ""
	if len(orders) >= int(params.PageSize)+1 {
		cursorData := ListOrderCursor{
			CreatedAt: orders[params.PageSize].OrderDate,
			ID:        orders[params.PageSize].ID,
		}
		cursor = pagination.EncodeCursor(cursorData)
		orders = orders[:params.PageSize]
	}

	return orders, cursor, nil
}
