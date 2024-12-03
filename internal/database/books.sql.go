// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: books.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const addUserToBook = `-- name: AddUserToBook :exec
UPDATE books
SET user_ids = array_append(user_ids, $2)
WHERE id = $1
`

type AddUserToBookParams struct {
	ID          uuid.UUID
	ArrayAppend interface{}
}

func (q *Queries) AddUserToBook(ctx context.Context, arg AddUserToBookParams) error {
	_, err := q.db.ExecContext(ctx, addUserToBook, arg.ID, arg.ArrayAppend)
	return err
}

const createBook = `-- name: CreateBook :one
INSERT INTO books (id, name, author, genre, user_ids)
VALUES (
    $1,
    $2,   -- name
    $3,   -- author
    $4,   -- genre
    $5    -- user_ids (UUID array)
)
RETURNING id, name, author, genre, user_ids
`

type CreateBookParams struct {
	ID      uuid.UUID
	Name    string
	Author  string
	Genre   string
	UserIds []uuid.UUID
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.ID,
		arg.Name,
		arg.Author,
		arg.Genre,
		pq.Array(arg.UserIds),
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Genre,
		pq.Array(&i.UserIds),
	)
	return i, err
}

const findBookByName = `-- name: FindBookByName :one
SELECT id, name, author, genre, user_ids
FROM books
WHERE name = $1
`

func (q *Queries) FindBookByName(ctx context.Context, name string) (Book, error) {
	row := q.db.QueryRowContext(ctx, findBookByName, name)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Author,
		&i.Genre,
		pq.Array(&i.UserIds),
	)
	return i, err
}