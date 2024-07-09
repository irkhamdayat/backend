-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_to_accesses(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL REFERENCES roles(id),
    access_id UUID NOT NULL REFERENCES accesses(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_to_accesses;
-- +goose StatementEnd