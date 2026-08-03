package local_store

import (
	"app/storage"
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestCreateConversation_GeneratesUUIDv7(t *testing.T) {
	store := NewTestStore(t)
	c, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "Test Conversation",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	id, err := uuid.Parse(c.ID)
	if err != nil {
		t.Fatalf("ID is not a valid UUID: %v", err)
	}
	if id.Version() != 7 {
		t.Errorf("expected UUIDv7, got version %d", id.Version())
	}
}

func TestCreateConversation_CreatesMessage(t *testing.T) {
	store := NewTestStore(t)
	c, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name:        "Test Conversation",
		MessageBody: "Hello, world!",
	})

	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}

	messages, err := store.ListMessagesByConversation(context.Background(), storage.ListMessagesByConversationParams{ConversationId: c.ID})

	if err != nil {
		t.Fatalf("ListMessagesByConversation failed: %v", err)
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(messages))
	}

	message := messages[0]

	if message.Body != "Hello, world!" {
		t.Fatalf("unexpected message body: %v", message.Body)
	}
}

func TestCreateConversation_RoundTrip(t *testing.T) {
	store := NewTestStore(t)

	created, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "My Conversation",
	})
	if err != nil {
		t.Fatalf("CreateConversation failed: %v", err)
	}
	if created.Name != "My Conversation" {
		t.Errorf("Name = %q, want %q", created.Name, "My Conversation")
	}
	if created.CreatedAt.IsZero() {
		t.Error("CreatedAt should be non-zero")
	}
	if created.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be non-zero")
	}
}

func TestListConversations(t *testing.T) {
	store := NewTestStore(t)

	_, err := store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "A",
	})
	if err != nil {
		t.Fatalf("CreateConversation A failed: %v", err)
	}
	_, err = store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name: "B",
	})
	if err != nil {
		t.Fatalf("CreateConversation B failed: %v", err)
	}

	list, err := store.ListConversations(context.Background())
	if err != nil {
		t.Fatalf("ListConversations failed: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("got %d conversations, want 2", len(list))
	}
}
