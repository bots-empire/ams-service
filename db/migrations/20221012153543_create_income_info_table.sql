-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ams.income_info
(
    user_id bigInt NOT NULL,
    bot_link text NOT NULL,
    bot_name text NOT NULL,
    income_source text NOT NULL,
    type_bot text NOT NULL,
);

CREATE UNIQUE INDEX info_uniq
ON ams.income_info(user_id, bot_link);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ams.income_info;
-- +goose StatementEnd