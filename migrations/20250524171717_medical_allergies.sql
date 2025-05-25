-- +goose Up
-- +goose StatementBegin
CREATE TABLE medical_diseases (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    disease_id BIGINT NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (disease_id) REFERENCES diseases(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE medical_diseases;
-- +goose StatementEnd
