-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS agents
(
    id                  UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    code                VARCHAR(50)              NOT NULL UNIQUE,
    email               VARCHAR(255)             NOT NULL UNIQUE,
    phone_number        VARCHAR(20),
    first_name          VARCHAR(255)             NOT NULL,
    last_name           VARCHAR(255)                      DEFAULT NULL,
    birth_date          DATE                     NOT NULL,
    birth_place         VARCHAR(255)             NOT NULL,
    photo               UUID REFERENCES upload_files (id),
    address             TEXT,
    location            VARCHAR(255),
    code_referral       VARCHAR(10),
    ktp_document        UUID                     NOT NULL REFERENCES upload_files (id),
    ktp_number          VARCHAR(20)              NOT NULL,
    npwp_document       UUID                     NOT NULL REFERENCES upload_files (id),
    npwp_number         VARCHAR(20)              NOT NULL,
    bank_account_number VARCHAR(20)              NOT NULL,
    bank_id             UUID REFERENCES banks (id),
    username            VARCHAR(50)              NOT NULL UNIQUE,
    pin                 TEXT                     NOT NULL,
    password            TEXT                     NOT NULL,
    status              VARCHAR(50)              NOT NULL,
    is_subscribe_news   BOOLEAN                           DEFAULT FALSE,
    approved_by         UUID REFERENCES admins (id) DEFAULT NULL,

    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_agents_username ON agents (username);
CREATE INDEX idx_agents_email ON agents (email);
CREATE INDEX idx_agents_code ON agents (code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_agents_username;
DROP INDEX IF EXISTS idx_agents_email;
DROP INDEX IF EXISTS idx_agents_code;
DROP TABLE IF EXISTS agents;
-- +goose StatementEnd