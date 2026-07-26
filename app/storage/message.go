package storage

import (
	"context"
	"fmt"
	"time"
)

type MessageType string

const (
	// Indicates a message sent by user
	MessageTypeUser MessageType = "user"
	// Indicates a message sent by llm
	MessageTypeLLM MessageType = "llm"
)

// Parses message type from string input. Returns error for invalid values
func ParseMessageType(input string) (MessageType, error) {
	switch MessageType(input) {
	case MessageTypeUser, MessageTypeLLM:
		return MessageType(input), nil
	default:
		return "", fmt.Errorf("invalid message type: %s", input)
	}
}

type Message struct {
	ID             string
	ConversationId string
	Body           string
	MessageType    MessageType
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateMessageParams struct {
	ConversationId string
	Body           string
	MessageType    MessageType
}

type ListMessagesByConversationParams struct {
	ConversationId string
}

type MessageStore interface {
	// Creates a message in a given conversation. Updates the updated_at field of the conversation
	CreateMessage(ctx context.Context, params CreateMessageParams) (Message, error)
	// Lists messages for a given conversation
	ListMessagesByConversation(ctx context.Context, params ListMessagesByConversationParams) ([]Message, error)
}
