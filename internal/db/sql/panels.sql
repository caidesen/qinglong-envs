-- name: GetPanelByID :one
SELECT *
FROM panels
WHERE id = ?
LIMIT 1;

-- name: ListPanels :many
SELECT *
FROM panels
ORDER BY id
LIMIT ? OFFSET ?;

-- name: CreatePanel :one
INSERT INTO panels (name, url, client_id, client_secret)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: DeletePanelByID :exec
DELETE
FROM panels
WHERE id = ?;

-- name: UpdatePanel :exec
UPDATE panels
SET name          = ?,
    url           = ?,
    client_id     = ?,
    client_secret = ?,
    updated_at    = CURRENT_TIMESTAMP
WHERE id = ?