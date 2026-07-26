package provider_service

import (
	"app/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"
)

type mockProviderStore struct {
	providers map[string]storage.Provider
	nextID    int
	createErr error
	getErr    error
	listErr   error
	updateErr error
	deleteErr error
}

func newMockStore() *mockProviderStore {
	return &mockProviderStore{providers: make(map[string]storage.Provider)}
}

func (m *mockProviderStore) CreateProvider(_ context.Context, params storage.CreateProviderParams) (storage.Provider, error) {
	if m.createErr != nil {
		return storage.Provider{}, m.createErr
	}
	m.nextID++
	id := fmt.Sprintf("test-uuid-%d", m.nextID)
	p := storage.Provider{
		ID:           id,
		Name:         params.Name,
		ProviderType: params.ProviderType,
		BaseURL:      params.BaseURL,
		APIKey:       params.APIKey,
		CreatedAt:    time.Unix(1000, 0),
		UpdatedAt:    time.Unix(1000, 0),
	}
	m.providers[p.ID] = p
	return p, nil
}

func (m *mockProviderStore) GetProvider(_ context.Context, id string) (storage.Provider, error) {
	if m.getErr != nil {
		return storage.Provider{}, m.getErr
	}
	p, ok := m.providers[id]
	if !ok {
		return storage.Provider{}, sql.ErrNoRows
	}
	return p, nil
}

func (m *mockProviderStore) ListProviders(_ context.Context) ([]storage.Provider, error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	result := make([]storage.Provider, 0, len(m.providers))
	for _, p := range m.providers {
		result = append(result, p)
	}
	return result, nil
}

func (m *mockProviderStore) UpdateProvider(_ context.Context, params storage.UpdateProviderParams) (storage.Provider, error) {
	if m.updateErr != nil {
		return storage.Provider{}, m.updateErr
	}
	p, ok := m.providers[params.ID]
	if !ok {
		return storage.Provider{}, sql.ErrNoRows
	}
	p.Name = params.Name
	p.ProviderType = params.ProviderType
	p.BaseURL = params.BaseURL
	p.APIKey = params.APIKey
	p.UpdatedAt = time.Unix(2000, 0)
	m.providers[params.ID] = p
	return p, nil
}

func (m *mockProviderStore) DeleteProvider(_ context.Context, id string) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	delete(m.providers, id)
	return nil
}

func TestCreateProvider(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	baseURL := "https://api.openai.com"
	apiKey := "sk-test"
	resp, err := svc.CreateProvider(ProviderInput{
		Name:         "Test",
		ProviderType: "openai",
		BaseURL:      &baseURL,
		APIKey:       &apiKey,
	})
	if err != nil {
		t.Fatalf("CreateProvider failed: %v", err)
	}
	if resp.ID != "test-uuid-1" {
		t.Errorf("ID = %q, want %q", resp.ID, "test-uuid-1")
	}
	if resp.Name != "Test" {
		t.Errorf("Name = %q, want %q", resp.Name, "Test")
	}
	if resp.ProviderType != "openai" {
		t.Errorf("ProviderType = %q, want %q", resp.ProviderType, "openai")
	}
}

func TestCreateProvider_Error(t *testing.T) {
	store := newMockStore()
	store.createErr = errors.New("db error")
	svc := New(store)

	_, err := svc.CreateProvider(ProviderInput{Name: "Test", ProviderType: "openai"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGetProvider(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	created, _ := svc.CreateProvider(ProviderInput{Name: "Test", ProviderType: "openai"})

	got, err := svc.GetProvider(created.ID)
	if err != nil {
		t.Fatalf("GetProvider failed: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("ID = %q, want %q", got.ID, created.ID)
	}
}

func TestGetProvider_NotFound(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	_, err := svc.GetProvider("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent provider, got nil")
	}
}

func TestListProviders(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	svc.CreateProvider(ProviderInput{Name: "A", ProviderType: "openai"})
	svc.CreateProvider(ProviderInput{Name: "B", ProviderType: "anthropic"})

	list, err := svc.ListProviders()
	if err != nil {
		t.Fatalf("ListProviders failed: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("got %d providers, want 2", len(list))
	}
}

func TestListProviders_Error(t *testing.T) {
	store := newMockStore()
	store.listErr = errors.New("db error")
	svc := New(store)

	_, err := svc.ListProviders()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestUpdateProvider(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	created, _ := svc.CreateProvider(ProviderInput{Name: "Original", ProviderType: "openai"})

	newBase := "https://new.api.com"
	updated, err := svc.UpdateProvider(created.ID, ProviderInput{
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
}

func TestUpdateProvider_Error(t *testing.T) {
	store := newMockStore()
	store.updateErr = errors.New("db error")
	svc := New(store)

	_, err := svc.UpdateProvider("id", ProviderInput{Name: "X", ProviderType: "openai"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestDeleteProvider(t *testing.T) {
	store := newMockStore()
	svc := New(store)

	created, _ := svc.CreateProvider(ProviderInput{Name: "ToDelete", ProviderType: "openai"})

	if err := svc.DeleteProvider(created.ID); err != nil {
		t.Fatalf("DeleteProvider failed: %v", err)
	}

	_, err := svc.GetProvider(created.ID)
	if err == nil {
		t.Error("expected error after deletion, got nil")
	}
}

func TestDeleteProvider_Error(t *testing.T) {
	store := newMockStore()
	store.deleteErr = errors.New("db error")
	svc := New(store)

	err := svc.DeleteProvider("id")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
