CREATE TABLE providers (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))),
    name TEXT NOT NULL,                          -- Friendly name (e.g., "My Claude API")
    provider_type TEXT NOT NULL,                 -- 'openai', 'anthropic', 'google', 'local'
    
    -- Connection Configuration
    base_url TEXT,                               -- Custom endpoint (Required for local, optional override for cloud providers)
    api_key TEXT,                                -- Encrypted or plain text API key (Null for local authless setups)
    
    -- Metadata
    created_at INTEGER NOT NULL DEFAULT (unixepoch()),
    updated_at INTEGER NOT NULL DEFAULT (unixepoch())
);

-- Indexing for quick lookups
CREATE INDEX idx_providers_type ON providers(provider_type);
