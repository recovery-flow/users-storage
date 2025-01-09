-- name: CreateUser :one
INSERT INTO users (
    id,
    username
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: UpdateFullUser :one
UPDATE users
SET
    username = $2,
    title = $3,
    avatar = $4,
    status = $5,
    bio = $6,
    city = $7,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateUsername :one
UPDATE users
SET
    username = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateAvatar :one
UPDATE users
SET
    avatar = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateBio :one
UPDATE users
SET
    bio = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateStatus :one
UPDATE users
SET
    status = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateTitle :one
UPDATE users
SET
    title = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateCity :one
UPDATE users
SET
    city = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: SearchUsers :many
SELECT id, username, title, status, avatar, bio, city,
FROM users
WHERE username ILIKE '%' || $1 || '%'
ORDER BY username
    LIMIT $2 OFFSET $3;
