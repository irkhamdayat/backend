-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS admins
(
    id                 UUID PRIMARY KEY                      DEFAULT uuid_generate_v4(),
    photo              UUID REFERENCES upload_files (id),
    first_name         VARCHAR(255)             NOT NULL,
    last_name          VARCHAR(255)                          DEFAULT NULL,
    username           VARCHAR(50)              NOT NULL UNIQUE,
    password           TEXT                     NOT NULL,
    email              VARCHAR(255)             NOT NULL UNIQUE,
    role_id            UUID                     NOT NULL REFERENCES roles (id),
    insurance_brand_id UUID REFERENCES insurance_brands (id) DEFAULT NULL,
    status             VARCHAR(50)              NOT NULL,

    created_at         TIMESTAMP WITH TIME ZONE NOT NULL     DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE NOT NULL     DEFAULT CURRENT_TIMESTAMP,
    deleted_at         TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_admins_username ON admins (username);
CREATE INDEX idx_admins_email ON admins (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_admins_username;
DROP INDEX IF EXISTS idx_admins_email;
DROP TABLE IF EXISTS admins;
-- +goose StatementEnd