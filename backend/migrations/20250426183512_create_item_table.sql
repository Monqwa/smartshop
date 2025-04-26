-- +goose Up
-- +goose StatementBegin
create table items
(
    id      uuid primary key,
    name    varchar,
    is_buy  boolean,
    list_id uuid
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table items;
-- +goose StatementEnd
