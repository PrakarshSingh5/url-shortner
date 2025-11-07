package database

import (
    "database/sql"
    "fmt"

    _ "modernc.org/sqlite"
)

const sqliteDriver = "sqlite"

// New opens a SQLite database located at the given path and ensures foreign keys support.
func New(path string) (*sql.DB, error) {
    dsn := fmt.Sprintf("file:%s?_pragma=busy_timeout(5000)&_pragma=foreign_keys(ON)", path)
    db, err := sql.Open(sqliteDriver, dsn)
    if err != nil {
        return nil, fmt.Errorf("open sqlite database: %w", err)
    }

    if err := db.Ping(); err != nil {
        _ = db.Close()
        return nil, fmt.Errorf("ping sqlite database: %w", err)
    }

    return db, nil
}

// Migrate creates the required tables if they do not exist.
func Migrate(db *sql.DB) error {
    const query = `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        slug TEXT UNIQUE NOT NULL,
        original_url TEXT NOT NULL,
        created_at TEXT NOT NULL DEFAULT (datetime('now'))
    )`

    if _, err := db.Exec(query); err != nil {
        return fmt.Errorf("migrate database: %w", err)
    }

    return nil
}
