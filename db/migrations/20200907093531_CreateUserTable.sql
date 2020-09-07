
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
    id serial not null primary key,
    name text not null
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS users;
