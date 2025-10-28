-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS folder();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS folder;
-- +goose StatementEnd
