-- name: CreateUser :one
INSERT INTO users (name, surname, email, password_hash, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: UpdateUserRole :one
UPDATE users SET role = $1 WHERE id = $2 RETURNING *;

-- name: CreateUserTg :one
INSERT INTO user_tgs (id, name, tg_username)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUserTgName :exec
UPDATE user_tgs
SET name = @name
WHERE id = @id;

-- name: GetUserTg :one
SELECT u.id, u.name AS user_name, u.unit_id, un.name AS unit_name
FROM user_tgs u
LEFT JOIN units un ON u.unit_id = un.id
WHERE u.id = $1;

-- name: GetUnits :many
SELECT * FROM units ORDER BY created_at DESC;

-- name: UpdateUserUnitID :exec
UPDATE user_tgs SET unit_id = @unit_id WHERE id = @user_id;

-- name: CreateApplication :one
INSERT INTO applications (text, unit_id, user_tg_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetApplications :many
SELECT 
    applications.id,
    applications.text,
    applications.status,
    sqlc.embed(user_tgs),
    sqlc.embed(units)
FROM applications
LEFT JOIN user_tgs ON applications.user_tg_id = user_tgs.id
LEFT JOIN units ON applications.unit_id = units.id
ORDER BY applications.created_at DESC;

-- name: GetApplicationsByUnitID :many
SELECT * FROM applications WHERE unit_id = $1 ORDER BY created_at DESC;

-- name: UpdateApplicationStatus :one
UPDATE applications SET status = $1 WHERE id = $2 RETURNING *;

-- name: DeleteApplication :exec
DELETE FROM applications WHERE id = $1;

-- name: CreateUnit :one
INSERT INTO units (name)
VALUES ($1)
RETURNING *;

-- name: DeleteUnit :exec
DELETE FROM units WHERE id = $1;

-- name: UpdateUnitName :one
UPDATE units SET name = $1 WHERE id = $2 RETURNING *;
