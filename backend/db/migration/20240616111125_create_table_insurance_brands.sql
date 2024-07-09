-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS insurance_brands
(
    id         UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    code       VARCHAR(20)              NOT NULL,
    icon       UUID REFERENCES upload_files (id),
    name       VARCHAR(50)              NOT NULL,
    category   VARCHAR(10)              NOT NULL,
    is_active  BOOLEAN,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_insurance_brands_name ON insurance_brands (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_insurance_brands_name;
DROP TABLE IF EXISTS insurance_brands;
-- +goose StatementEnd