-- name: CreateUser :one
INSERT INTO users (name, surname, email, password_hash, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUserTg :one
INSERT INTO users_tg (id, name, tg_username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUserTgName :exec
UPDATE users_tg
SET name = @name
WHERE id = @id;

-- name: GetUserTg :one
SELECT u.id, u.name AS user_name, u.unit_id, un.name AS unit_name
FROM users_tg u
LEFT JOIN units un ON u.unit_id = un.id
WHERE u.id = $1;

-- name: GetUnits :many
SELECT * FROM units;

-- name: UpdateUserUnitID :exec
UPDATE users_tg SET unit_id = @unit_id WHERE id = @user_id;

-- name: CreateApplication :one
INSERT INTO applications (text, unit_id, user_tg_id)
VALUES ($1, $2, $3)
RETURNING *;
