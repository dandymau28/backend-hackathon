-- +goose Up
-- +goose StatementBegin
CREATE TABLE medical_allergies (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    allergy_id BIGINT NOT NULL,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (allergy_id) REFERENCES allergies(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE medical_allergies;
-- +goose StatementEnd
