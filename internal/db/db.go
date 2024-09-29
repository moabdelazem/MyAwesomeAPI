package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxOpenConnection int, maxIdleConnection int, maxIdleTime time.Duration) (*sql.DB, error) {
	// Create a new connection to the database
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	// Verfiy Connection To Database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// Set the maximum number of open connections
	db.SetMaxOpenConns(maxOpenConnection)

	// Set the maximum number of idle connections
	db.SetMaxIdleConns(maxIdleConnection)

	// Set the maximum idle time for a connection
	db.SetConnMaxIdleTime(maxIdleTime)

	// Ping the database to ensure the connection is working
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
