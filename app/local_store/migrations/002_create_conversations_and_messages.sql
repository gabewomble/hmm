CREATE TABLE conversations (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,

    -- Metadata
    created_at INTEGER NOT NULL DEFAULT (unixepoch()),
    updated_at INTEGER NOT NULL DEFAULT (unixepoch())
);

CREATE INDEX idx_conversations_name ON conversations(name);

CREATE TABLE messages (
    id TEXT PRIMARY KEY,
    conversation_id TEXT NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    seq INTEGER NOT NULL,
    body TEXT NOT NULL,
    message_type TEXT NOT NULL CHECK (message_type IN ('user', 'llm')),

    -- Metadata
    created_at INTEGER NOT NULL DEFAULT (unixepoch()),
    updated_at INTEGER NOT NULL DEFAULT (unixepoch())
);

-- Composite index for fetching messages in order per conversation
CREATE INDEX idx_messages_conversation ON messages(conversation_id, seq);
