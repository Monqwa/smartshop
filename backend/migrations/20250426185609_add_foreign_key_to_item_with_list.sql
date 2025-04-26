-- +goose Up
-- +goose StatementBegin
ALTER TABLE items
    ADD CONSTRAINT fk_items_list
        FOREIGN KEY (list_id)
            REFERENCES list(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items
DROP CONSTRAINT fk_items_list;
-- +goose StatementEnd
