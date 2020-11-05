// When sqlc supports custom models,
// this file will be generated

package gen

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.xixo.com/api/services/identity/domain/users"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(first_name, last_name, email, phone_number, account_id)
  VALUES ($1, $2, $3, $4, $5) RETURNING user_id
`

type CreateUserParams struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber sql.NullString
	AccountID   uuid.UUID
	Roles       []uuid.UUID
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (users.User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.AccountID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		pq.Array(arg.Roles),
	)
	var i users.User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE account_id = $1 AND user_id = $2
`

type DeleteUserParams struct {
	AccountID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser, arg.AccountID, arg.UserID)
	return err
}

const denyRoleFromUser = `-- name: DenyRoleFromUser :exec
DELETE FROM users_roles WHERE user_id = $1 AND role_id = $2
`

type DenyRoleFromUserParams struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (q *Queries) DenyRoleFromUser(ctx context.Context, arg DenyRoleFromUserParams) error {
	_, err := q.db.ExecContext(ctx, denyRoleFromUser, arg.UserID, arg.RoleID)
	return err
}

const findUserByID = `-- name: FindUserByID :one
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles WHERE account_id = $1 AND user_id = $2
`

type FindUserByIDParams struct {
	AccountID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) FindUserByID(ctx context.Context, arg FindUserByIDParams) (users.User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, arg.AccountID, arg.UserID)
	var i users.User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}

const findUsers = `-- name: FindUsers :many
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles
    WHERE account_id = $1 ORDER BY created_at DESC, user_id DESC LIMIT $2
`

type FindUsersParams struct {
	AccountID uuid.UUID
	Limit     int32
}

func (q *Queries) FindUsers(ctx context.Context, arg FindUsersParams) ([]users.User, error) {
	rows, err := q.db.QueryContext(ctx, findUsers, arg.AccountID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []users.User
	for rows.Next() {
		var i users.User
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.Roles),
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

const findUsersCursor = `-- name: FindUsersCursor :many
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles
    WHERE account_id = $1 AND (
      created_at <= $2 AND (
        created_at < $2 OR user_id < $3
      )
    )
    ORDER BY created_at DESC, user_id DESC LIMIT $4
`

type FindUsersCursorParams struct {
	AccountID uuid.UUID
	CreatedAt time.Time
	UserID    uuid.UUID
	Limit     int32
}

func (q *Queries) FindUsersCursor(ctx context.Context, arg FindUsersCursorParams) ([]users.User, error) {
	rows, err := q.db.QueryContext(ctx, findUsersCursor,
		arg.AccountID,
		arg.CreatedAt,
		arg.UserID,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []users.User
	for rows.Next() {
		var i users.User
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.Roles),
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

const grantAccountAdminRoleToUser = `-- name: GrantAccountAdminRoleToUser :exec
INSERT INTO users_roles(user_id, role_id) VALUES ($1, (
  SELECT role_id FROM roles WHERE account_admin
)) ON CONFLICT DO NOTHING
`

func (q *Queries) GrantAccountAdminRoleToUser(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, grantAccountAdminRoleToUser, userID)
	return err
}

const grantRoleToUser = `-- name: GrantRoleToUser :exec
INSERT INTO users_roles(user_id, role_id) VALUES ($1, $2)
`

type GrantRoleToUserParams struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (q *Queries) GrantRoleToUser(ctx context.Context, arg GrantRoleToUserParams) error {
	_, err := q.db.ExecContext(ctx, grantRoleToUser, arg.UserID, arg.RoleID)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users_with_roles SET
	first_name = CASE WHEN $1 THEN $2 ELSE first_name END,
	last_name = CASE WHEN $3 THEN $4 ELSE last_name END,
	email = CASE WHEN $5 THEN $6 ELSE email END,
	phone_number = CASE WHEN $7 THEN $8 ELSE phone_number END,
	roles = CASE WHEN $9 THEN $10 ELSE roles END
WHERE user_id = $11
	RETURNING user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles;
`

type UpdateUserParams struct {
	FirstName               string
	ShouldUpdateFirstName   bool
	LastName                string
	ShouldUpdateLastName    bool
	Email                   string
	ShouldUpdateEmail       bool
	PhoneNumber             sql.NullString
	ShouldUpdatePhoneNumber bool
	Roles                   []string
	ShouldUpdateRoles       bool
	UserID                  uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (users.User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.ShouldUpdateFirstName,
		arg.LastName,
		arg.ShouldUpdateLastName,
		arg.Email,
		arg.ShouldUpdateEmail,
		arg.PhoneNumber,
		arg.ShouldUpdatePhoneNumber,
		pq.Array(arg.Roles),
		arg.ShouldUpdateRoles,
		arg.UserID,
	)
	var i users.User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}
