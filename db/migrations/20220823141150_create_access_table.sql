-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS ams;

CREATE TABLE IF NOT EXISTS ams.accesses
(
    user_id     bigint  NOT NULL,
    code        text    NOT NULL,
    additional  text[],

    user_name       text NOT NULL,
    user_first_name text NOT NULL,
    user_last_name  text NOT NULL
);

CREATE UNIQUE INDEX access_uniq
ON ams.accesses(user_id, code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA ams CASCADE;
-- +goose StatementEnd
