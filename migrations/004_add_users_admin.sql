-- +migrate Up
INSERT INTO users (
    username,
    password,
    created_at,
    created_by
)
SELECT
    'admin',
    '$2a$12$v/A7kyU3cDnW9T53S44MUOoEFBO78DPB7bZHEsvPpGNV3YA7rRPbe',
    NOW(),
    'system'
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE username = 'admin'
);