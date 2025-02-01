package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

func NewSqliteConn(storagePath string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	stmt, err := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		topic TEXT NOT NULL,
		description TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return conn, nil
}
