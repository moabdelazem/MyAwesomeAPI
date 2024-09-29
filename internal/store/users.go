package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// User is a struct that defines the fields of a user
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
}

// UserStore is a struct that defines the methods to interact with the users table
type UserStore struct {
	db *sql.DB
}

// Create is a method on the UserStore that creates a new user in the database
func (us *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	// Execute the query
	err := us.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(&user.ID, &user.CreatedAt)

	// Handle any errors
	if err != nil {
		return err
	}

	return nil
}

// Get All The Users In The Database
func (us *UserStore) GetUsers(ctx context.Context) ([]User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
	`

	rows, err := us.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
