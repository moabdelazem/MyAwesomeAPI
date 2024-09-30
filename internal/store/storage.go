package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

// Posts interface
type Posts interface {
	Create(context.Context, *Post) error
}

// Users interface
type Users interface {
	Create(context.Context, *sql.Tx, *User) error
	GetUsers(context.Context) ([]User, error)
	GetUserByID(context.Context, uuid.UUID) (*User, error)
	CreateUser(context.Context, *User) error
	GetUserByUsername(context.Context, string) (*User, error)
}

// Storage Store
type Storage struct {
	Posts Posts
	Users Users
}

// NewStorage creates a new storage
// and returns a pointer to it
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db: db},
		Users: &UserStore{db: db},
	}
}

func WithTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	// Begin a new transaction
	tx, err := db.BeginTx(ctx, nil)
	// Handle any errors
	if err != nil {
		return err
	}

	// Defer a rollback in case of a panic
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
