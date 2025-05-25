-- +goose Up
-- +goose StatementBegin
CREATE TABLE calorie_counts (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    datetime timestamptz NOT NULL,
    calories INT NOT NULL DEFAULT 0,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE calorie_counts;
-- +goose StatementEnd
