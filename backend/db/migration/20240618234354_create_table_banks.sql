-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banks
(
    id         UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    code       VARCHAR(20)              NOT NULL,
    icon       UUID REFERENCES upload_files (id),
    name       VARCHAR(50)              NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_banks_name ON banks (name);
CREATE INDEX idx_banks_code ON banks (code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_banks_name;
DROP INDEX IF EXISTS idx_banks_code;
DROP TABLE IF EXISTS banks;
-- +goose StatementEnd