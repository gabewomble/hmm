CREATE TABLE conversations (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,

    -- Metadata
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_conversations_name ON conversations(name);

CREATE TABLE messages (
    id TEXT PRIMARY KEY,
    conversation_id TEXT NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    message_type TEXT NOT NULL CHECK (message_type IN ('user', 'llm')),

    -- Metadata
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index for fetching messages by conversation
CREATE INDEX idx_messages_conversation ON messages(conversation_id);
