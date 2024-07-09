
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notification_histories
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    is_read BOOL DEFAULT FALSE,
    image VARCHAR DEFAULT NULL,
    notification_type VARCHAR(30) NOT NULL,
    action_type VARCHAR(30) NOT NULL,
    additional_data JSONB,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX ON notification_histories(created_at);
CREATE INDEX ON notification_histories(notification_type);
CREATE INDEX ON notification_histories(action_type);

CREATE TABLE IF NOT EXISTS translate_notification_histories(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    notification_history_id UUID NOT NULL REFERENCES notification_histories(id),
    language VARCHAR(4) not null,
    headline TEXT NOT NULL,
    message TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop indexes for notification_histories
DROP INDEX IF EXISTS notification_histories_created_at_idx;
DROP INDEX IF EXISTS notification_histories_notification_type_idx;
DROP INDEX IF EXISTS notification_histories_action_type_idx;

-- Drop table translate_notification_histories
DROP TABLE IF EXISTS translate_notification_histories;

-- Drop table notification_histories
DROP TABLE IF EXISTS notification_histories;
-- +goose StatementEnd
