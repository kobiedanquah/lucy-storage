-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS folders(
    id uuid NOT NULL,
    parent_id uuid,
    user_id uuid NOT NULL,
    name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    last_modified TIMESTAMP DEFAULT now() NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES folders (id) ON DELETE CASCADE,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS folders;
-- +goose StatementEnd
