package local_store

import (
	"app/storage"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Creates a new conversation
func (s *LocalStore) CreateConversation(ctx context.Context, params storage.CreateConversationParams) (storage.Conversation, error) {
	conversationId, err := uuid.NewV7()
	if err != nil {
		return storage.Conversation{}, err
	}

	now := time.Now()

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return storage.Conversation{}, err
	}

	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	result, err := qtx.CreateConversation(ctx, CreateConversationParams{
		ID:        conversationId.String(),
		Name:      params.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return storage.Conversation{}, err
	}

	if params.MessageBody != "" {
		messageId, err := uuid.NewV7()
		if err != nil {
			return storage.Conversation{}, err
		}

		if _, err := qtx.CreateMessage(ctx, CreateMessageParams{
			ID:             messageId.String(),
			ConversationID: conversationId.String(),
			Body:           params.MessageBody,
			MessageType:    string(storage.MessageTypeUser),
			CreatedAt:      now,
			UpdatedAt:      now,
		}); err != nil {
			return storage.Conversation{}, err
		}
	}

	err = tx.Commit()

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
