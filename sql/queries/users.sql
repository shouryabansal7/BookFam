-- name: CreateUser :one
INSERT INTO users (id, name, email, password)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: AddBookIDToUser :exec
UPDATE users
SET book_ids = array_append(book_ids, $2)
WHERE id = $1;