-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_roles AS ENUM ('USER', 'ADMIN');
CREATE TABLE users (
    uuid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email TEXT,
    first_name TEXT,
    last_name TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT false,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW(),
    roles user_roles ARRAY NOT NULL,
    tsv tsvector
);
CREATE INDEX users_tsv_idx ON users USING gin(tsv);
INSERT INTO users(email, first_name, last_name, roles, tsv) VALUES ('testuser@goboilerplate.io', 'John', 'Smith', '{USER, ADMIN}', setweight(to_tsvector('John'), 'A') || setweight(to_tsvector('Smith'), 'B'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd