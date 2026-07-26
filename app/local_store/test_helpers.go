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

	store := &LocalStore{db: db, queries: New(db)}

	if err := store.runMigrations(false); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}

	return store
}
