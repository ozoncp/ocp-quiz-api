-- +goose Up
-- +goose StatementBegin
CREATE TABLE quiz
(
    id           SERIAL PRIMARY KEY,
    classroom_id bigint not null,
    user_id      bigint not null,
    link         text   not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quiz;
-- +goose StatementEnd
