-- +goose Up
ALTER TABLE users
ADD COLUMN book_ids UUID[];

-- +goose Down
ALTER TABLE users
DROP COLUMN book_ids;