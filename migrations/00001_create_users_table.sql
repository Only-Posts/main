-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id       SERIAL PRIMARY KEY,
    username CHAR(100) UNIQUE NOT NULL,
    email    CHAR(200) UNIQUE NOT NULL,
    password CHAR(128)        NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
