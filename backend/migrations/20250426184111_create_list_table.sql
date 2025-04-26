-- +goose Up
-- +goose StatementBegin
create table list (
    id uuid primary key,
    title varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table list;
-- +goose StatementEnd
