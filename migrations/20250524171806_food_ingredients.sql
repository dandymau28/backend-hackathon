-- +goose Up
-- +goose StatementBegin
CREATE TABLE food_ingredients (
    id SERIAL PRIMARY KEY,
    food_id BIGINT NOT NULL,
    name varchar(255) NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (food_id) REFERENCES merchant_foods(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE food_ingredients;
-- +goose StatementEnd
