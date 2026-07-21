package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/adrg/xdg"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var database *sql.DB

func Init() error {
	dbDir := filepath.Join(xdg.DataHome, "hmm")
	dbPath := filepath.Join(dbDir, "hmm.db")

	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	var err error
	database, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err := database.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	if err := runMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Printf("Database initialized at %s", dbPath)
	return nil
}

func Close() error {
	if database != nil {
		return database.Close()
	}
	return nil
}

func DB() *sql.DB {
	return database
}

func runMigrations() error {
	if _, err := database.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (name TEXT PRIMARY KEY, applied_at INTEGER NOT NULL DEFAULT (unixepoch()))`); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	entries, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations: %w", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		var applied bool
		if err := database.QueryRow("SELECT 1 FROM schema_migrations WHERE name = ?", name).Scan(&applied); err == nil {
			continue
		}

		content, err := fs.ReadFile(migrationsFS, filepath.Join("migrations", name))
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", name, err)
		}

		if _, err := database.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", name, err)
		}

		if _, err := database.Exec("INSERT INTO schema_migrations (name) VALUES (?)", name); err != nil {
			return fmt.Errorf("failed to record migration %s: %w", name, err)
		}

		log.Printf("Applied migration: %s", name)
	}

	return nil
}
