package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (

	rxEmail    = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	rxUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{0,17}$")

	// ErrUserNotFound used when the user wasnt found
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidEmail used when the email is not valid
	ErrInvalidEmail = errors.New("invalid email")

	// ErrInvalidUsername used when the email is not valid
	ErrInvalidUsername = errors.New("invalid username")

	// ErrEmailTaken used when new data with same email already inputed
	ErrEmailTaken = errors.New("email already taken")

	// ErrUsernameTaken used when new data with same username already inputed
	ErrUsernameTaken = errors.New("username already taken")
)

// User Model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// CreateUser inserts a user in the database
func (s *Service) CreateUser(ctx context.Context, email, username string) error {

	email = strings.TrimSpace(email)
	if !rxEmail.MatchString(email) {
		return ErrInvalidEmail
	}

	username = strings.TrimSpace(username)
	if !rxUsername.MatchString(username) {
		return ErrInvalidUsername
	}

	query := "INSERT INTO users (email, username) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, email, username)
	unique := isUniqueViolation(err)

	if unique && strings.Contains(err.Error(), "email") {
		return ErrEmailTaken
	}

	if unique && strings.Contains(err.Error(), "username") {
		return ErrUsernameTaken
	}

	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
 }