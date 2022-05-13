package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type DB struct {
	url string
}

// InitDB returns a new database instance.
func InitDB(cfg string) *DB {
	return &DB{url: cfg}
}

// Connect returns a database connection.
func (db *DB) Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), db.url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	// Check database connection
	if err := conn.Ping(context.Background()); err != nil {
		return nil, errors.Wrap(err, "ping database failed")
	}

	return conn, nil
}
