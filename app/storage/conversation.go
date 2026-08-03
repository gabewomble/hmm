package storage

import (
	"context"
	"time"
)

type Conversation struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateConversationParams struct {
	// Name of the conversation
	Name string
	// Body of message which creates conversation
	MessageBody string
}

type ConversationStore interface {
	// Creates a new conversation
	CreateConversation(ctx context.Context, params CreateConversationParams) (Conversation, error)
	// Lists conversations
	ListConversations(ctx context.Context) ([]Conversation, error)
}
