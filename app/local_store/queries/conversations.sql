-- name: CreateConversation :one
INSERT INTO conversations (
    id,
    name,
    created_at,
    updated_at
) VALUES (
    ?,
    ?,
    ?,
    ?
) RETURNING *;

-- name: ListConversations :many
SELECT * FROM conversations ORDER BY updated_at DESC;

-- name: UpdateConversationUpdatedAt :exec
UPDATE conversations SET updated_at = ? WHERE id = ?;

