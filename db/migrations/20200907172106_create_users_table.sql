
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS users (
   id          SERIAL         PRIMARY KEY,
   name        VARCHAR(191)   NOT NULL,
   email       VARCHAR(191)   NOT NULL,
   password    VARCHAR(191)   NOT NULL,
   token       VARCHAR(191)   NOT NULL,
   created_at  DATETIME       NOT NULL,
   updated_at  DATETIME       NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS users;
