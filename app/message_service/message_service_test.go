package message_service

import (
	"app/storage"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type mockMessageStore struct {
	messages  map[string]storage.Message
	nextID    int
	createErr error
	listErr   error
}

func newMockMessageStore() *mockMessageStore {
	return &mockMessageStore{messages: make(map[string]storage.Message)}
}

func (m *mockMessageStore) CreateMessage(_ context.Context, params storage.CreateMessageParams) (storage.Message, error) {
	if m.createErr != nil {
		return storage.Message{}, m.createErr
	}
	m.nextID++
	id := fmt.Sprintf("test-uuid-%d", m.nextID)
	now := time.Now()
	msg := storage.Message{
		ID:             id,
		ConversationId: params.ConversationId,
		Body:           params.Body,
		MessageType:    params.MessageType,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	m.messages[msg.ID] = msg
	return msg, nil
}

func (m *mockMessageStore) ListMessagesByConversation(_ context.Context, params storage.ListMessagesByConversationParams) ([]storage.Message, error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	result := make([]storage.Message, 0)
	for _, msg := range m.messages {
		if msg.ConversationId == params.ConversationId {
			result = append(result, msg)
		}
	}
	return result, nil
}

func TestCreateMessage(t *testing.T) {
	store := newMockMessageStore()
	svc := New(store)

	resp, err := svc.CreateMessage(CreateMessageRequest{
		ConversationId: "conv-123",
		Body:           "Hello world",
	})
	if err != nil {
		t.Fatalf("CreateMessage failed: %v", err)
	}
	if resp.ID != "test-uuid-1" {
		t.Errorf("ID = %q, want %q", resp.ID, "test-uuid-1")
	}
	if resp.ConversationId != "conv-123" {
		t.Errorf("ConversationId = %q, want %q", resp.ConversationId, "conv-123")
	}
	if resp.Body != "Hello world" {
		t.Errorf("Body = %q, want %q", resp.Body, "Hello world")
	}
	if resp.MessageType != storage.MessageTypeUser {
		t.Errorf("MessageType = %q, want %q", resp.MessageType, storage.MessageTypeUser)
	}
	if resp.CreatedAt.IsZero() {
		t.Error("CreatedAt should be non-zero")
	}
	if resp.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be non-zero")
	}
}

func TestCreateMessage_Error(t *testing.T) {
	store := newMockMessageStore()
	store.createErr = errors.New("db error")
	svc := New(store)

	_, err := svc.CreateMessage(CreateMessageRequest{
		ConversationId: "conv-123",
		Body:           "Hello",
	})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestListMessagesByConversation(t *testing.T) {
	store := newMockMessageStore()
	svc := New(store)

	svc.CreateMessage(CreateMessageRequest{
		ConversationId: "conv-1",
		Body:           "First",
	})
	svc.CreateMessage(CreateMessageRequest{
		ConversationId: "conv-1",
		Body:           "Second",
	})
	svc.CreateMessage(CreateMessageRequest{
		ConversationId: "conv-2",
		Body:           "Other conversation",
	})

	list, err := svc.ListMessagesByConversation(ListMessagesByConversationRequest{
		ConversationId: "conv-1",
	})
	if err != nil {
		t.Fatalf("ListMessagesByConversation failed: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("got %d messages, want 2", len(list))
	}
}

func TestListMessagesByConversation_Error(t *testing.T) {
	store := newMockMessageStore()
	store.listErr = errors.New("db error")
	svc := New(store)

	_, err := svc.ListMessagesByConversation(ListMessagesByConversationRequest{
		ConversationId: "conv-1",
	})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestListMessagesByConversation_Empty(t *testing.T) {
	store := newMockMessageStore()
	svc := New(store)

	list, err := svc.ListMessagesByConversation(ListMessagesByConversationRequest{
		ConversationId: "conv-empty",
	})
	if err != nil {
		t.Fatalf("ListMessagesByConversation failed: %v", err)
	}
	if len(list) != 0 {
		t.Errorf("expected 0 messages, got %d", len(list))
	}
}
