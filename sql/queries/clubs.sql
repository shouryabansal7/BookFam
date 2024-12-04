-- name: CreateClubs :one
INSERT INTO clubs (id, name, genre, user_ids)
VALUES (
    $1,
    $2,   -- name
    $3,   -- genre
    $4    -- user_ids (UUID array)
)
RETURNING *;

-- name: AddUserClub :exec
UPDATE clubs
SET user_ids = array_append(user_ids, $2)
WHERE id = $1;

-- name: RemoveUserFromClub :exec
UPDATE clubs
SET user_ids = array_remove(user_ids, $2)
WHERE id = $1;

-- name: GetClubs :many
SELECT * FROM clubs;