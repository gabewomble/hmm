-- name: CreateProvider :one
INSERT INTO providers (
    name,
    provider_type,
    base_url,
    api_key,
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

-- name: GetProvider :one
SELECT * FROM providers WHERE id = ?;

-- name: ListProviders :many
SELECT * FROM providers ORDER BY created_at DESC;

-- name: UpdateProvider :one
UPDATE providers SET
    name = ?,
    provider_type = ?,
    base_url = ?,
    api_key = ?,
    updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteProvider :exec
DELETE FROM providers WHERE id = ?;
