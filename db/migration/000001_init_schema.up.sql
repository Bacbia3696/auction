CREATE TABLE users (
    id bigserial PRIMARY KEY,
    user_name varchar NOT NULL,
    password varchar NOT NULL,
    full_name varchar NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);
