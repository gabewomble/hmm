package local_store

import (
	"app/storage"
	"context"
	"time"

	"github.com/google/uuid"
)

// Creates a new conversation
func (s *LocalStore) CreateConversation(ctx context.Context, params storage.CreateConversationParams) (storage.Conversation, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return storage.Conversation{}, err
	}

	now := time.Now()
	result, err := s.queries.CreateConversation(ctx, CreateConversationParams{
		ID:        id.String(),
		Name:      params.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return storage.Conversation{}, err
	}

	return toStorageConversation(result), nil
}

// Lists conversations
func (s *LocalStore) ListConversations(ctx context.Context) ([]storage.Conversation, error) {
	result, err := s.queries.ListConversations(ctx)

	if err != nil {
		return nil, err
	}

	conversations := make([]storage.Conversation, len(result))

	for idx, conversation := range result {
		conversations[idx] = toStorageConversation(conversation)
	}

	return conversations, nil
}

// Converts the db Conversation struct to the storage Conversation struct
func toStorageConversation(c Conversation) storage.Conversation {
	return storage.Conversation{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
