package store

import (
	"context"
	"database/sql"
)

type Posts interface {
	Create(context.Context, *Post) error
}

type Users interface {
	Create(context.Context, *User) error
	GetUsers(context.Context) ([]User, error)
}

type Storage struct {
	Posts Posts
	Users Users
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db: db},
		Users: &UserStore{db: db},
	}
}
