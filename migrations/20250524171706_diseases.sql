-- +goose Up
-- +goose StatementBegin
CREATE TABLE diseases (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE diseases;
-- +goose StatementEnd
