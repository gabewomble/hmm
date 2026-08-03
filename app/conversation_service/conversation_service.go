package conversation_service

import (
	"app/storage"
	"context"
	"time"
)

type ConversationService struct {
	store storage.ConversationStore
}

func New(store storage.ConversationStore) *ConversationService {
	return &ConversationService{
		store: store,
	}
}

type CreateConversationRequest struct {
	Name        string `json:"name"`
	MessageBody string `json:"messageBody"`
}

type ConversationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Create a new conversation
func (s *ConversationService) CreateConversation(input CreateConversationRequest) (ConversationResponse, error) {
	result, err := s.store.CreateConversation(context.Background(), storage.CreateConversationParams{
		Name:        input.Name,
		MessageBody: input.MessageBody,
	})

	if err != nil {
		return ConversationResponse{}, err
	}

	return toResponse(result), nil
}

type DeleteConversationRequest struct {
	ConversationId string `json:"conversationId"`
}

// Deletes a conversation
func (s *ConversationService) DeleteConversation(input DeleteConversationRequest) error {
	return s.store.DeleteConversation(context.Background(), storage.DeleteConversationParams{
		ConversationId: input.ConversationId,
	})
}

// Lists conversations
func (s *ConversationService) ListConversations() ([]ConversationResponse, error) {
	result, err := s.store.ListConversations(context.Background())

	if err != nil {
		return nil, err
	}

	conversations := make([]ConversationResponse, len(result))

	for idx, conversation := range result {
		conversations[idx] = toResponse(conversation)
	}

	return conversations, nil
}

func toResponse(c storage.Conversation) ConversationResponse {
	return ConversationResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
