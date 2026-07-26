package message_service

import (
	"app/storage"
	"context"
	"time"
)

type MessageService struct {
	store storage.MessageStore
}

// Returns a new instance of a message service
func New(store storage.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

type CreateMessageRequest struct {
	ConversationId string `json:"conversationId"` // Conversation ID message belongs to
	Body           string `json:"body"`           // Body of message
}

type MessageResponse struct {
	ID             string              `json:"id"`             // Message ID
	ConversationId string              `json:"conversationId"` // Conversation ID message belongs to
	Body           string              `json:"body"`           // Body of message
	MessageType    storage.MessageType `json:"messageType"`    // Type of message ('user' | 'llm')
	CreatedAt      time.Time           `json:"createdAt"`      // When the message was created
	UpdatedAt      time.Time           `json:"updatedAt"`      // When the message was updated
}

// Create a new message in the given conversation
func (s *MessageService) CreateMessage(input CreateMessageRequest) (MessageResponse, error) {
	result, err := s.store.CreateMessage(context.Background(), storage.CreateMessageParams{
		ConversationId: input.ConversationId,
		Body:           input.Body,
		MessageType:    storage.MessageTypeUser,
	})

	if err != nil {
		return MessageResponse{}, err
	}

	return toMessageResponse(result), nil
}

type ListMessagesByConversationRequest struct {
	ConversationId string `json:"conversationId"` // Conversation ID message belongs to
}

// Lists messages for a given conversation
func (s *MessageService) ListMessagesByConversation(input ListMessagesByConversationRequest) ([]MessageResponse, error) {
	result, err := s.store.ListMessagesByConversation(context.Background(), storage.ListMessagesByConversationParams{
		ConversationId: input.ConversationId,
	})

	if err != nil {
		return nil, err
	}

	messages := make([]MessageResponse, len(result))

	for idx, message := range result {
		messages[idx] = toMessageResponse(message)
	}

	return messages, nil
}

func toMessageResponse(m storage.Message) MessageResponse {
	return MessageResponse{
		ID:             m.ID,
		ConversationId: m.ConversationId,
		Body:           m.Body,
		MessageType:    m.MessageType,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}
