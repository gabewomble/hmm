package local_store

import (
	"app/storage"
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

func newTestStore(t *testing.T) *LocalStore {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	if _, err := db.Exec(`CREATE TABLE providers (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		provider_type TEXT NOT NULL,
		base_url TEXT,
		api_key TEXT,
		created_at INTEGER NOT NULL DEFAULT (unixepoch()),
		updated_at INTEGER NOT NULL DEFAULT (unixepoch())
	)`); err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	return &LocalStore{db: db, queries: New(db)}
}

func TestCreateProvider_GeneratesUUIDv7(t *testing.T) {
	store := newTestStore(t)
	p, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name:         "Test",
		ProviderType: "openai",
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	id, err := uuid.Parse(p.ID)
	if err != nil {
		t.Fatalf("ID is not a valid UUID: %v", err)
	}
	if id.Version() != 7 {
		t.Errorf("expected UUIDv7, got version %d", id.Version())
	}
}

func TestCreateProvider_RoundTrip(t *testing.T) {
	store := newTestStore(t)
	baseURL := "https://api.openai.com"
	apiKey := "sk-test"

	created, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name:         "My Provider",
		ProviderType: "openai",
		BaseURL:      &baseURL,
		APIKey:       &apiKey,
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}
	if created.Name != "My Provider" {
		t.Errorf("Name = %q, want %q", created.Name, "My Provider")
	}
	if created.ProviderType != "openai" {
		t.Errorf("ProviderType = %q, want %q", created.ProviderType, "openai")
	}
	if created.BaseURL == nil || *created.BaseURL != baseURL {
		t.Errorf("BaseURL = %v, want %q", created.BaseURL, baseURL)
	}
	if created.APIKey == nil || *created.APIKey != apiKey {
		t.Errorf("APIKey = %v, want %q", created.APIKey, apiKey)
	}
	if created.CreatedAt == 0 {
		t.Error("CreatedAt should be non-zero")
	}

	got, err := store.GetProvider(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("GetProvider failed: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("GetProvider ID = %q, want %q", got.ID, created.ID)
	}
	if got.Name != created.Name {
		t.Errorf("GetProvider Name = %q, want %q", got.Name, created.Name)
	}
}

func TestListProviders(t *testing.T) {
	store := newTestStore(t)

	_, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name: "A", ProviderType: "openai",
	})
	if err != nil {
		t.Fatalf("CreateProvider A failed: %v", err)
	}
	_, err = store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name: "B", ProviderType: "anthropic",
	})
	if err != nil {
		t.Fatalf("CreateProvider B failed: %v", err)
	}

	list, err := store.ListProviders(context.Background())
	if err != nil {
		t.Fatalf("ListProviders failed: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("got %d providers, want 2", len(list))
	}
}

func TestUpdateProvider(t *testing.T) {
	store := newTestStore(t)
	newBase := "https://new.api.com"

	created, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name: "Original", ProviderType: "openai",
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	updated, err := store.UpdateProvider(context.Background(), storage.UpdateProviderParams{
		ID:           created.ID,
		Name:         "Updated",
		ProviderType: "anthropic",
		BaseURL:      &newBase,
	})
	if err != nil {
		t.Fatalf("UpdateProvider failed: %v", err)
	}
	if updated.Name != "Updated" {
		t.Errorf("Name = %q, want %q", updated.Name, "Updated")
	}
	if updated.ProviderType != "anthropic" {
		t.Errorf("ProviderType = %q, want %q", updated.ProviderType, "anthropic")
	}
	if updated.BaseURL == nil || *updated.BaseURL != newBase {
		t.Errorf("BaseURL = %v, want %q", updated.BaseURL, newBase)
	}
}

func TestDeleteProvider(t *testing.T) {
	store := newTestStore(t)

	created, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name: "ToDelete", ProviderType: "openai",
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}

	if err := store.DeleteProvider(context.Background(), created.ID); err != nil {
		t.Fatalf("DeleteProvider failed: %v", err)
	}

	_, err = store.GetProvider(context.Background(), created.ID)
	if err == nil {
		t.Error("expected error getting deleted provider, got nil")
	}
}

func TestCreateProvider_NullableFields(t *testing.T) {
	store := newTestStore(t)

	p, err := store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name:         "No Optionals",
		ProviderType: "local",
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}
	if p.BaseURL != nil {
		t.Errorf("BaseURL = %v, want nil", p.BaseURL)
	}
	if p.APIKey != nil {
		t.Errorf("APIKey = %v, want nil", p.APIKey)
	}
}
