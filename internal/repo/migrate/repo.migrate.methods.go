package migrate

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/huseinnashr/pforder-backend/internal/domain"
	"github.com/huseinnashr/pforder-backend/internal/pkg/tracer"
)

func (r *Repo) CSVToDB(ctx context.Context, dbTx domain.ISQLDatabaseTx, params domain.CSVToDBParam) error {
	ctx, span := tracer.Start(ctx, "repo.CSVToDB")
	defer span.End()

	file, err := os.Open(params.Filepath)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(file)
	isHeader := true
	count := int64(1)
	placeholders := []string{}
	data := []interface{}{}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if isHeader == true {
			isHeader = false
			continue
		}
		placeholder := []string{}
		for i, field := range record {
			placeholder = append(placeholder, fmt.Sprintf("$%d", count))

			columnName := ""
			if i < len(params.Columns) {
				columnName = params.Columns[i]
			}

			if transformer := params.Transform[columnName]; transformer != nil {
				transformedField, err := transformer(field)
				if err != nil {
					return err
				}
				data = append(data, transformedField)
			} else {
				data = append(data, field)
			}

			count++
		}

		placeholders = append(placeholders, "("+strings.Join(placeholder, ",")+")")
	}

	query := fmt.Sprintf(
		`INSERT INTO %s(%s) VALUES %s ON CONFLICT DO NOTHING`,
		params.TableName,
		strings.Join(params.Columns, ","),
		strings.Join(placeholders, ","),
	)
	if _, err = dbTx.ExecContext(ctx, query, data...); err != nil {
		return err
	}

	return nil
}
