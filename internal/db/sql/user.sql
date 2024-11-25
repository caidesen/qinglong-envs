-- name: GetUserById :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = ?
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT ? OFFSET ?;

-- name: CreateUser :one
INSERT INTO users (username, password, wx_pusher_uid)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET username    = ?,
    password    = ?,
    wx_pusher_uid = ?
WHERE id = ?