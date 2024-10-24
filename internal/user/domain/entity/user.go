package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedBy string
	CreatedAt time.Time
}

// NewUser return a new user.
// It can only accept password with length up to 72 bytes.
func NewUser(email, password string) (User, error) {
	hashedPwd, err := hashPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{
		Email:     email,
		Password:  hashedPwd,
		CreatedBy: email,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// VerifyPassword compare user's hashed password with the plain password given.
func (u User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
