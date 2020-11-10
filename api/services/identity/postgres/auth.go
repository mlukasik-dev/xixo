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
var _ auth.Repository = (*Repository)(nil)

// CheckUsersPassword .
func (r *Repository) CheckUsersPassword(ctx context.Context, accountID uuid.UUID, email, plainPassword string) (bool, error) {
	hash, err := r.q.FindUsersPassword(ctx, gen.FindUsersPasswordParams{
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

// UpdateUsersPassword .
func (r *Repository) UpdateUsersPassword(ctx context.Context, accountID uuid.UUID, email, plainPassword string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	err = r.q.UpdateUsersPassword(ctx, gen.UpdateUsersPasswordParams{
		AccountID: accountID,
		Email:     email,
		Password:  sql.NullString{String: string(hash), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

// CheckAdminsPassword .
func (r *Repository) CheckAdminsPassword(ctx context.Context, email, plainPassword string) (bool, error) {
	hash, err := r.q.FindAdminsPassword(ctx, email)
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

// UpdateAdminsPassword .
func (r *Repository) UpdateAdminsPassword(ctx context.Context, email, plainPassword string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return err
	}
	err = r.q.UpdateAdminsPassword(ctx, gen.UpdateAdminsPasswordParams{
		Email:    email,
		Password: sql.NullString{String: string(hash), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

// FindUserInfoByEmail .
func (r *Repository) FindUserInfoByEmail(ctx context.Context, accountID uuid.UUID, email string) (*auth.UserInfo, error) {
	info, err := r.q.FindUserInfoByEmail(ctx, gen.FindUserInfoByEmailParams{
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

// FindAdminInfoByEmail .
func (r *Repository) FindAdminInfoByEmail(ctx context.Context, email string) (*auth.AdminInfo, error) {
	info, err := r.q.FindAdminInfoByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &info, nil
}
