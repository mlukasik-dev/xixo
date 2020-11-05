package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/api/services/identity/postgres/gen"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// verify interface compliance
var _ auth.Repository = (*repo)(nil)

func (r *repo) CheckUsersPassword(accountID uuid.UUID, email, plainPassword string) (bool, error) {
	hash, err := r.q.FindUsersPassword(context.Background(), gen.FindUsersPasswordParams{
		AccountID: accountID, Email: email,
	})
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

func (r *repo) UpdateUsersPassword(accountID uuid.UUID, email, plainPassword string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	err = r.q.UpdateUsersPassword(context.Background(), gen.UpdateUsersPasswordParams{
		AccountID: accountID,
		Email:     email,
		Password:  sql.NullString{String: string(hash), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CheckAdminsPassword(email, plainPassword string) (bool, error) {
	hash, err := r.q.FindAdminsPassword(context.Background(), email)
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

func (r *repo) UpdateAdminsPassword(email, plainPassword string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	err = r.q.UpdateAdminsPassword(context.Background(), gen.UpdateAdminsPasswordParams{
		Email:    email,
		Password: sql.NullString{String: string(hash), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindUserInfoByEmail(accountID uuid.UUID, email string) (*auth.UserInfo, error) {
	info, err := r.q.FindUserInfoByEmail(context.Background(), gen.FindUserInfoByEmailParams{
		AccountID: accountID,
		Email:     email,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *repo) FindAdminInfoByEmail(email string) (*auth.AdminInfo, error) {
	info, err := r.q.FindAdminInfoByEmail(context.Background(), email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}
