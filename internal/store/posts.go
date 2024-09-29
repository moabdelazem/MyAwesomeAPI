package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Post is a struct that defines the fields of a post
type Post struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Tags      []string  `json:"tags"`
}

// PostsStore is a struct that defines the methods to interact with the posts table
type PostsStore struct {
	db *sql.DB
}

// Create is a method on the PostsStore that creates a new post in the database
func (ps *PostsStore) Create(ctx context.Context, post *Post) error {
	// Define the query
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	// Execute the query
	err := ps.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(
			post.Tags,
		),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	// Handle any errors
	if err != nil {
		return err
	}

	return nil
}
