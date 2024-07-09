-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accesses(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    label varchar NOT NULL,
    value varchar NOT NULL,
    section varchar NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accesses;
-- +goose StatementEnd