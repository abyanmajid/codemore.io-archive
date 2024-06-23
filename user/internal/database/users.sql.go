// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, email, password, roles, xp, is_banned, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, username, email, password, roles, xp, is_banned, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	Roles     []string
	Xp        int32
	IsBanned  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		pq.Array(arg.Roles),
		arg.Xp,
		arg.IsBanned,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		pq.Array(&i.Roles),
		&i.Xp,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			pq.Array(&i.Roles),
			&i.Xp,
			&i.IsBanned,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		pq.Array(&i.Roles),
		&i.Xp,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, email, password, roles, xp, is_banned, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		pq.Array(&i.Roles),
		&i.Xp,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3, password = $4, roles = $5, xp = $6, is_banned = $7, updated_at = NOW()
WHERE id = $1
RETURNING id, username, email, password, roles, xp, is_banned, created_at, updated_at
`

type UpdateUserParams struct {
	ID       uuid.UUID
	Username string
	Email    string
	Password string
	Roles    []string
	Xp       int32
	IsBanned bool
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		pq.Array(arg.Roles),
		arg.Xp,
		arg.IsBanned,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		pq.Array(&i.Roles),
		&i.Xp,
		&i.IsBanned,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
