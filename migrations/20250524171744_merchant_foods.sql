-- +goose Up
-- +goose StatementBegin
CREATE TABLE merchant_foods (
    id SERIAL PRIMARY KEY,
    merchant_id BIGINT NOT NULL,
    name varchar(255) NOT NULL,
    price int NOT NULL DEFAULT 0,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE merchant_foods;
-- +goose StatementEnd
