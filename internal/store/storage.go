package store

import (
	"context"
	"database/sql"
)

// Posts interface
type Posts interface {
	Create(context.Context, *Post) error
}

// Users interface
type Users interface {
	Create(context.Context, *User) error
	GetUsers(context.Context) ([]User, error)
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
