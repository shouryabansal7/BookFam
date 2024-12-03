-- name: CreateBook :one
INSERT INTO books (id, name, author, genre, user_ids)
VALUES (
    $1,
    $2,   -- name
    $3,   -- author
    $4,   -- genre
    $5    -- user_ids (UUID array)
)
RETURNING *;

-- name: FindBookByName :one
SELECT id, name, author, genre, user_ids
FROM books
WHERE name = $1;

-- name: AddUserToBook :exec
UPDATE books
SET user_ids = array_append(user_ids, $2)
WHERE id = $1;

-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBookById :one
SELECT * FROM books WHERE id = $1;