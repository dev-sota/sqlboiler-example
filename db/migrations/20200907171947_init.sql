
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
SET CHARSET utf8mb4;
ALTER DATABASE sqlboiler_example DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

