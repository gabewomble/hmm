package local_store

import (
	"app/storage"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateMessage_GeneratesUUIDv7(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	msg, err := store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Hello",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage failed: %v", err)
	}

	id, err := uuid.Parse(msg.ID)
	if err != nil {
		t.Fatalf("ID is not a valid UUID: %v", err)
	}
	if id.Version() != 7 {
		t.Errorf("expected UUIDv7, got version %d", id.Version())
	}
}

func TestCreateMessage_RoundTrip(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	msg, err := store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Hello world",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage failed: %v", err)
	}

	if msg.ConversationId != conv.ID {
		t.Errorf("ConversationId = %q, want %q", msg.ConversationId, conv.ID)
	}
	if msg.Body != "Hello world" {
		t.Errorf("Body = %q, want %q", msg.Body, "Hello world")
	}
	if msg.MessageType != storage.MessageTypeUser {
		t.Errorf("MessageType = %q, want %q", msg.MessageType, storage.MessageTypeUser)
	}
	if msg.CreatedAt.IsZero() {
		t.Error("CreatedAt should be non-zero")
	}
	if msg.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be non-zero")
	}
}

func TestCreateMessage_UpdatesConversationTimestamp(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	originalUpdatedAt := conv.UpdatedAt

	time.Sleep(10 * time.Millisecond)

	_, err = store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Hello",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage failed: %v", err)
	}

	list, err := store.ListConversations(context.Background())
	if err != nil {
		t.Fatalf("ListConversations failed: %v", err)
	}

	if len(list) != 1 {
		t.Fatalf("expected 1 conversation, got %d", len(list))
	}

	if !list[0].UpdatedAt.After(originalUpdatedAt) {
		t.Error("conversation updated_at should have been updated after creating a message")
	}
}

func TestCreateMessage_InvalidType(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	_, err = store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Hello",
		MessageType:    "invalid",
	})
	if err == nil {
		t.Error("expected error for invalid message type, got nil")
	}
}

func TestListMessagesByConversation_Ordering(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	msg1, err := store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "First",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage 1 failed: %v", err)
	}

	time.Sleep(10 * time.Millisecond)

	msg2, err := store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Second",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage 2 failed: %v", err)
	}

	time.Sleep(10 * time.Millisecond)

	msg3, err := store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: conv.ID,
		Body:           "Third",
		MessageType:    storage.MessageTypeUser,
	})
	if err != nil {
		t.Fatalf("CreateMessage 3 failed: %v", err)
	}

	messages, err := store.ListMessagesByConversation(context.Background(), storage.ListMessagesByConversationParams{
		ConversationId: conv.ID,
	})
	if err != nil {
		t.Fatalf("ListMessagesByConversation failed: %v", err)
	}

	if len(messages) != 3 {
		t.Fatalf("expected 3 messages, got %d", len(messages))
	}

	if messages[0].ID != msg1.ID {
		t.Errorf("first message ID = %q, want %q", messages[0].ID, msg1.ID)
	}
	if messages[1].ID != msg2.ID {
		t.Errorf("second message ID = %q, want %q", messages[1].ID, msg2.ID)
	}
	if messages[2].ID != msg3.ID {
		t.Errorf("third message ID = %q, want %q", messages[2].ID, msg3.ID)
	}
}

func TestListMessagesByConversation_Empty(t *testing.T) {
	store := NewTestStore(t)

	conv, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	messages, err := store.ListMessagesByConversation(context.Background(), storage.ListMessagesByConversationParams{
		ConversationId: conv.ID,
	})
	if err != nil {
		t.Fatalf("ListMessagesByConversation failed: %v", err)
	}

	if len(messages) != 0 {
		t.Errorf("expected 0 messages, got %d", len(messages))
	}
}
