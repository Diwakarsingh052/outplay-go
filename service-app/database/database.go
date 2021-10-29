package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Config is the required properties to use the database.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	TLSMode  string
}

// Open knows how to open a database connection based on the configuration.
func Open(cfg Config) (*sql.DB, error) {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.TLSMode)
	return sql.Open("postgres", conn)
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *sql.DB) error {

	// Run a simple query to determine connectivity. The db has a "Ping" method
	// but it can false-positive when it was previously able to talk to the
	// database but the database has since gone away. Running this query forces a
	// round trip to the database.
	const q = `SELECT true`
	var tmp bool
	return db.QueryRowContext(ctx, q).Scan(&tmp)
}
