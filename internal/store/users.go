package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
	ErrNotFound          = errors.New("user not found")
)

// User is a struct that defines the fields of a user
type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  Password  `json:"-"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	IsActive  bool      `json:"is_active"`
	// RoleID    int64     `json:"role_id"`
}

// Password is a struct that defines the fields of a password
type Password struct {
	Text *string
	Hash []byte
}

// Set is a method on the Password struct that hashes a password
// and sets the Text and Hash fields
func (p *Password) Set(text string) error {
	hash, err := hashPassword(text)
	if err != nil {
		return err
	}

	p.Text = &text
	p.Hash = hash

	return nil
}

// hashPassword is a helper function that hashes a password
// using bcrypt and returns the hashed password
func hashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

// UserStore is a struct that defines the methods to interact with the users table
type UserStore struct {
	db *sql.DB
}

// Create is a method on the UserStore that creates a new user in the database
func (us *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
		INSERT INTO users (id, username, password, email)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at
	`

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Execute the query
	err := tx.QueryRowContext(
		ctx,
		query,
		uuid.New(),
		user.Username,
		user.Password.Hash,
		user.Email,
	).Scan(&user.ID, &user.CreatedAt)

	// Handle any errors
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

// Get User By ID
func (us *UserStore) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
		WHERE id = $1
	`

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Execute the query
	row := us.db.QueryRowContext(ctx, query, id)

	// Scan the row into the User struct
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Get All The Users In The Database
// !TESTING PURPOSES ONLY!
// !DELETE THIS FUNCTION IN PRODUCTION
func (us *UserStore) GetUsers(ctx context.Context) ([]User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
	`

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Execute the query
	rows, err := us.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
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

// Get User By Email
func (us *UserStore) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, username, email, created_at
		FROM users
		WHERE email = $1
	`

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Execute the query
	row := us.db.QueryRowContext(ctx, query, email)

	// Scan the row into the User struct
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		}
	}

	return &user, nil
}

// CreateUser creates a new user in the database
func (us *UserStore) CreateUser(ctx context.Context, user *User) error {
	return WithTx(us.db, ctx, func(tx *sql.Tx) error {
		if err := us.Create(ctx, tx, user); err != nil {
			return err
		}

		return nil
	})

}
