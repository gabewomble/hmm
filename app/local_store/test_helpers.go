package local_store

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func NewTestStore(t *testing.T) *LocalStore {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	// Enable foreign_keys
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		t.Fatalf("failed to enable foreign_keys: %v", err)
	}

	store := &LocalStore{db: db, queries: New(db)}

	if err := store.runMigrations(false); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}

	return store
}
