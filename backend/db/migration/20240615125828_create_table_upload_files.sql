-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS upload_files(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    file_name VARCHAR(256) NOT NULL,
    upload_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX ON upload_files (upload_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS upload_files;
-- +goose StatementEnd
