package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/auth"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// verify interface compliance
var _ auth.Repository = (*repo)(nil)

func (r *repo) CheckUsersPassword(accountID, email, plainPassword string) (bool, error) {
	const query = `
		SELECT password FROM users WHERE account_id = $1 AND email = $2
	`
	var hash sql.NullString
	err := r.db.GetContext(context.Background(), &hash, query, accountID, email)
	if errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return false, err
	}
	if !hash.Valid {
		return false, auth.ErrNoPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash.String), []byte(plainPassword))
	return err == nil, nil
}

func (r *repo) SetUsersPassword(accountID, email, plainPassword string) error {
	const query = `
		UPDATE users SET password = $1 WHERE account_id = $2 AND email = $3
	`
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(context.Background(), query, hash, accountID, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CheckAdminsPassword(email, plainPassword string) (bool, error) {
	const query = `
		SELECT password FROM admins WHERE email = $1
	`
	var hash sql.NullString
	err := r.db.GetContext(context.Background(), &hash, query, email)
	if errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return false, err
	}
	if !hash.Valid {
		return false, auth.ErrNoPassword
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash.String), []byte(plainPassword))
	return err == nil, nil
}

func (r *repo) SetAdminsPassword(email, plainPassword string) error {
	const query = `
		UPDATE admins SET password = $1 WHERE email = $2
	`
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(context.Background(), query, hash, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindUserInfoByEmail(accountID, email string) (*auth.UserInfo, error) {
	const query = `
		SELECT user_id, registered, roles
			FROM users_with_roles
				WHERE account_id = $1 AND email = $2 
	`
	var info auth.UserInfo
	var nullRoles []sql.NullString
	err := r.db.QueryRowContext(context.Background(), query, accountID, email).
		Scan(
			&info.ID, &info.Registered, pq.Array(&nullRoles),
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	info.RoleIDs = str.NullStringsToStrings(nullRoles)
	return &info, nil
}

func (r *repo) FindAdminInfoByEmail(email string) (*auth.AdminInfo, error) {
	const query = `
		SELECT admin_id, registered, roles
			FROM admins_with_roles WHERE email = $1
	`
	var info auth.AdminInfo
	var nullRoles []sql.NullString
	err := r.db.QueryRowContext(context.Background(), query, email).
		Scan(
			&info.ID, &info.Registered, pq.Array(&nullRoles),
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	info.RoleIDs = str.NullStringsToStrings(nullRoles)
	return &info, nil
}
