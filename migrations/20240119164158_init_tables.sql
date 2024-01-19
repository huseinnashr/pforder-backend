-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS public.orders (
    id bigserial PRIMARY KEY,
    created_at timestamp NOT NULL,
    order_name text NOT NULL,
    customer_id text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.order_items (
    id bigserial PRIMARY KEY,
    order_id bigint NOT NULL,
    price_per_unit bigint NOT NULL,
    quantity int NOT NULL,
    product text NOT NULL
);

CREATE TABLE IF NOT EXISTS public.deliveries (
    id bigserial PRIMARY KEY,
    order_item_id bigint NOT NULL,
    delivered_quantity int NOT NULL
);

CREATE TABLE IF NOT EXISTS public.customers (
    user_id text PRIMARY KEY,
    login text NOT NULL UNIQUE,
    password text NOT NULL,
    name text NOT NULL,
    company_id bigint NOT NULL,
    credit_cards text[] NOT NULL
);

CREATE TABLE IF NOT EXISTS public.customer_companies (
    company_id bigserial PRIMARY KEY,
    company_name text NOT NULL
);

ALTER TABLE public.orders ADD CONSTRAINT fk_customers FOREIGN KEY (customer_id) REFERENCES customers(user_id);
ALTER TABLE public.order_items ADD CONSTRAINT fk_orders FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE public.deliveries ADD CONSTRAINT fk_order_items FOREIGN KEY (order_item_id) REFERENCES order_items(id);
ALTER TABLE public.customers ADD CONSTRAINT fk_customer_companies FOREIGN KEY (company_id) REFERENCES customer_companies(company_id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE public.customers DROP CONSTRAINT IF EXISTS fk_customer_companies;
ALTER TABLE public.deliveries DROP CONSTRAINT IF EXISTS fk_order_items;
ALTER TABLE public.order_items DROP CONSTRAINT IF EXISTS fk_orders;
ALTER TABLE public.orders DROP CONSTRAINT IF EXISTS fk_customers;

DROP TABLE IF EXISTS public.customer_companies;
DROP TABLE IF EXISTS public.customers;
DROP TABLE IF EXISTS public.deliveries;
DROP TABLE IF EXISTS public.order_items;
DROP TABLE IF EXISTS public.orders;
