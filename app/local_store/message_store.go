package local_store

import (
	"app/storage"
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *LocalStore) CreateMessage(ctx context.Context, params storage.CreateMessageParams) (storage.Message, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return storage.Message{}, err
	}

	now := time.Now()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return storage.Message{}, err
	}
	defer tx.Rollback()

	q := s.queries.WithTx(tx)

	result, err := q.CreateMessage(ctx, CreateMessageParams{
		ID:             id.String(),
		ConversationID: params.ConversationId,
		Body:           params.Body,
		MessageType:    string(params.MessageType),
		CreatedAt:      now,
		UpdatedAt:      now,
	})
	if err != nil {
		return storage.Message{}, err
	}

	if err := q.UpdateConversationUpdatedAt(ctx, UpdateConversationUpdatedAtParams{
		UpdatedAt: now,
		ID:        params.ConversationId,
	}); err != nil {
		return storage.Message{}, err
	}

	if err := tx.Commit(); err != nil {
		return storage.Message{}, err
	}

	return toStorageMessage(result), nil
}

func (s *LocalStore) ListMessagesByConversation(ctx context.Context, params storage.ListMessagesByConversationParams) ([]storage.Message, error) {
	result, err := s.queries.ListMessagesByConversation(ctx, params.ConversationId)

	if err != nil {
		return nil, err
	}

	messages := make([]storage.Message, len(result))

	for idx, message := range result {
		messages[idx] = toStorageMessage(message)
	}

	return messages, nil
}

func toStorageMessage(m Message) storage.Message {
	return storage.Message{
		ID:             m.ID,
		ConversationId: m.ConversationID,
		Body:           m.Body,
		MessageType:    storage.MessageType(m.MessageType),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}
