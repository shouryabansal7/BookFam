-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,                     -- Auto-incrementing primary key for Books
    name VARCHAR(255) NOT NULL,                 -- Name of the book
    author VARCHAR(255) NOT NULL,               -- Author of the book
    genre VARCHAR(255) NOT NULL,                -- Genre of the book
    user_ids UUID[]                           -- Array of UUIDs representing users currently reading the book
);

-- +goose Down
DROP TABLE books;