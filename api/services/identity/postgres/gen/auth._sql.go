// When sqlc supports custom models,
// this file will be generated

package gen

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.xixo.com/api/services/identity/domain/auth"
)

const findAdminInfoByEmail = `-- name: FindAdminInfoByEmail :one
SELECT admin_id, registered, roles
  FROM admins_with_roles WHERE email = $1
`

func (q *Queries) FindAdminInfoByEmail(ctx context.Context, email string) (auth.AdminInfo, error) {
	row := q.db.QueryRowContext(ctx, findAdminInfoByEmail, email)
	var i auth.AdminInfo
	err := row.Scan(&i.ID, &i.Registered, pq.Array(&i.Roles))
	return i, err
}

const findAdminsPassword = `-- name: FindAdminsPassword :one
SELECT password FROM admins WHERE email = $1
`

func (q *Queries) FindAdminsPassword(ctx context.Context, email string) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, findAdminsPassword, email)
	var password sql.NullString
	err := row.Scan(&password)
	return password, err
}

const findUserInfoByEmail = `-- name: FindUserInfoByEmail :one
SELECT user_id, registered, roles
  FROM users_with_roles
    WHERE account_id = $1 AND email = $2
`

type FindUserInfoByEmailParams struct {
	AccountID uuid.UUID
	Email     string
}

func (q *Queries) FindUserInfoByEmail(ctx context.Context, arg FindUserInfoByEmailParams) (auth.UserInfo, error) {
	row := q.db.QueryRowContext(ctx, findUserInfoByEmail, arg.AccountID, arg.Email)
	var i auth.UserInfo
	err := row.Scan(&i.ID, &i.Registered, pq.Array(&i.Roles))
	return i, err
}

const findUsersPassword = `-- name: FindUsersPassword :one
SELECT password FROM users WHERE account_id = $1 AND email = $2
`

type FindUsersPasswordParams struct {
	AccountID uuid.UUID
	Email     string
}

func (q *Queries) FindUsersPassword(ctx context.Context, arg FindUsersPasswordParams) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, findUsersPassword, arg.AccountID, arg.Email)
	var password sql.NullString
	err := row.Scan(&password)
	return password, err
}

const updateAdminsPassword = `-- name: UpdateAdminsPassword :exec
UPDATE admins SET password = $1 WHERE email = $2
`

type UpdateAdminsPasswordParams struct {
	Password sql.NullString
	Email    string
}

func (q *Queries) UpdateAdminsPassword(ctx context.Context, arg UpdateAdminsPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateAdminsPassword, arg.Password, arg.Email)
	return err
}

const updateUsersPassword = `-- name: UpdateUsersPassword :exec
UPDATE users SET password = $1 WHERE account_id = $2 AND email = $3
`

type UpdateUsersPasswordParams struct {
	Password  sql.NullString
	AccountID uuid.UUID
	Email     string
}

func (q *Queries) UpdateUsersPassword(ctx context.Context, arg UpdateUsersPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUsersPassword, arg.Password, arg.AccountID, arg.Email)
	return err
}
