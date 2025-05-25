-- +goose Up
-- +goose StatementBegin
CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    ratings int NOT NULL DEFAULT 0,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE merchants;
-- +goose StatementEnd
