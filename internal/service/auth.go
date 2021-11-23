package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// TokenLifespan until tokens are valid
	TokenLifespan = time.Hour * 24 * 14
)

// LoginOutput Response
type LoginOutput struct {
	Token string
	ExpiresAt time.Time
	AuthUser User
}

// Login insecurerly
func (s *Service) Login(ctx context.Context, email string) (LoginOutput, error) {

	var out LoginOutput

	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return out, ErrInvalidEmail
	}

	query := "SELECT id, username FROM users WHERE email =$1"
	err := s.db.QueryRowContext(ctx, query, email).Scan(&out.AuthUser.ID, &out.AuthUser.Username)
	if err == sql.ErrNoRows {
		return out, ErrUserNotFound
	}

	if err != nil {
		return out, fmt.Errorf("could not query select user: %v", err)
	}

	out.Token,err = s.codec.EncodeToString(strconv.FormatInt(out.AuthUser.ID, 10))
	if err != nil {
		return out, fmt.Errorf("could not create token: %v", err)
	}

	out.ExpiresAt = time.Now().Add(TokenLifespan)

	return out, nil

}