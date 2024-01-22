-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX order_search ON orders USING gin (order_name gin_trgm_ops);
CREATE INDEX order_range_filter ON orders(created_at, id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP INDEX IF EXISTS order_range_filter;
DROP INDEX IF EXISTS order_search;
DROP EXTENSION IF EXISTS pg_trgm;