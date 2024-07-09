-- +goose Up
-- +goose StatementBegin
INSERT INTO accesses (id, "label", value, "section", created_at, updated_at, deleted_at)
VALUES ('89deab56-b10d-41ff-bd36-ea7f8963777e'::UUID, 'Execute Decision Withdrawal Agent', 'EXECUTE_WITHDRAW_AGENT',
        'WITHDRAW_MENU', '2024-05-16 17:00:00.000', '2024-05-16 17:00:00.000', NULL),
       ('e4930e9c-afce-4d49-a7e1-42fdd9a6ec83'::UUID, 'View Withdrawal Agent', 'VIEW_WITHDRAW_AGENT', 'WITHDRAW_MENU',
        '2024-05-16 17:00:00.000', '2024-05-16 17:00:00.000', NULL);

INSERT INTO roles (id, "name", created_at, updated_at, deleted_at)
VALUES ('c853dea6-9b5b-40df-aab4-6a7665731e24'::UUID, 'Finance', '2024-05-16 17:00:00.000', '2024-05-16 17:00:00.000',
        NULL);

INSERT INTO role_to_accesses (id, "role_id", access_id)
VALUES ('1dd009e1-103b-4576-ac32-fd079de26989'::UUID, 'c853dea6-9b5b-40df-aab4-6a7665731e24',
        '89deab56-b10d-41ff-bd36-ea7f8963777e'),
       ('18ed1304-5f0c-4436-b4e5-1a6fa8f7743b'::UUID, 'c853dea6-9b5b-40df-aab4-6a7665731e24',
        'e4930e9c-afce-4d49-a7e1-42fdd9a6ec83');


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DELETE
FROM role_to_accesses
WHERE id IN (
             '1dd009e1-103b-4576-ac32-fd079de26989'::UUID,
             '18ed1304-5f0c-4436-b4e5-1a6fa8f7743b'::UUID
    );

DELETE
FROM accesses
WHERE id IN (
             '89deab56-b10d-41ff-bd36-ea7f8963777e'::UUID,
             'e4930e9c-afce-4d49-a7e1-42fdd9a6ec83'::UUID
    );

DELETE
FROM roles
WHERE id IN (
    'c853dea6-9b5b-40df-aab4-6a7665731e24'::UUID
    );
-- +goose StatementEnd
