package conversation_service

import (
	"app/storage"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type mockConversationStore struct {
	conversations map[string]storage.Conversation
	nextID        int
	createErr     error
	listErr       error
	deleteErr     error
}

func newMockConversationStore() *mockConversationStore {
	return &mockConversationStore{conversations: make(map[string]storage.Conversation)}
}

func (m *mockConversationStore) CreateConversation(_ context.Context, params storage.CreateConversationParams) (storage.Conversation, error) {
	if m.createErr != nil {
		return storage.Conversation{}, m.createErr
	}
	m.nextID++
	id := fmt.Sprintf("test-uuid-%d", m.nextID)
	now := time.Now()
	c := storage.Conversation{
		ID:        id,
		Name:      params.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	m.conversations[c.ID] = c
	return c, nil
}

func (m *mockConversationStore) DeleteConversation(_ context.Context, params storage.DeleteConversationParams) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}

	return nil
}

func (m *mockConversationStore) ListConversations(_ context.Context) ([]storage.Conversation, error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	result := make([]storage.Conversation, 0, len(m.conversations))
	for _, c := range m.conversations {
		result = append(result, c)
	}
	return result, nil
}

func TestCreateConversation(t *testing.T) {
	store := newMockConversationStore()
	svc := New(store)

	resp, err := svc.CreateConversation(CreateConversationRequest{
		Name:        "Test Conversation",
		MessageBody: "Hello",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}
	if resp.ID != "test-uuid-1" {
		t.Errorf("ID = %q, want %q", resp.ID, "test-uuid-1")
	}
	if resp.Name != "Test Conversation" {
		t.Errorf("Name = %q, want %q", resp.Name, "Test Conversation")
	}
	if resp.CreatedAt.IsZero() {
		t.Error("CreatedAt should be non-zero")
	}
	if resp.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be non-zero")
	}
}

func TestCreateConversation_Error(t *testing.T) {
	store := newMockConversationStore()
	store.createErr = errors.New("db error")
	svc := New(store)

	_, err := svc.CreateConversation(CreateConversationRequest{Name: "Test"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestListConversations(t *testing.T) {
	store := newMockConversationStore()
	svc := New(store)

	svc.CreateConversation(CreateConversationRequest{Name: "A"})
	svc.CreateConversation(CreateConversationRequest{Name: "B"})

	list, err := svc.ListConversations()
	if err != nil {
		t.Fatalf("ListConversations failed: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("got %d conversations, want 2", len(list))
	}
}

func TestListConversations_Error(t *testing.T) {
	store := newMockConversationStore()
	store.listErr = errors.New("db error")
	svc := New(store)

	_, err := svc.ListConversations()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestListConversations_Empty(t *testing.T) {
	store := newMockConversationStore()
	svc := New(store)

	list, err := svc.ListConversations()
	if err != nil {
		t.Fatalf("ListConversations failed: %v", err)
	}
	if len(list) != 0 {
		t.Errorf("expected 0 conversations, got %d", len(list))
	}
}

func TestDeleteConversation(t *testing.T) {
	store := newMockConversationStore()
	svc := New(store)

	err := svc.DeleteConversation(DeleteConversationRequest{ConversationId: "conversation-id"})
	if err != nil {
		t.Fatalf("DeleteConversation failed: %v", err)
	}
}

func TestDeleteConversation_Error(t *testing.T) {
	store := newMockConversationStore()
	store.deleteErr = errors.New("db error")
	svc := New(store)

	err := svc.DeleteConversation(DeleteConversationRequest{ConversationId: "conversation-id"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
