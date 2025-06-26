-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext NOT NULL UNIQUE, -- Stores text data exactly as is inputted, but comparisons against the data are always case-insensitive.
    password_hash bytea NOT NULL, -- Binary string.
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
