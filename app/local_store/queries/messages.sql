-- name: CreateMessage :one
INSERT INTO messages (
    id,
    conversation_id,
    body,
    message_type,
    created_at,
    updated_at
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
) RETURNING *;

-- name: ListMessagesByConversation :many
SELECT * FROM messages WHERE conversation_id = ? ORDER BY created_at ASC;
