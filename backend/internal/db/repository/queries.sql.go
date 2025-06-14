// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createApplication = `-- name: CreateApplication :one
INSERT INTO applications (text, unit_id, user_tg_id)
VALUES ($1, $2, $3)
RETURNING id, text, status, unit_id, user_tg_id, created_at, updated_at
`

type CreateApplicationParams struct {
	Text     string      `json:"text"`
	UnitID   uuid.UUID   `json:"unit_id"`
	UserTgID pgtype.Int8 `json:"user_tg_id"`
}

func (q *Queries) CreateApplication(ctx context.Context, arg CreateApplicationParams) (Application, error) {
	row := q.db.QueryRow(ctx, createApplication, arg.Text, arg.UnitID, arg.UserTgID)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Status,
		&i.UnitID,
		&i.UserTgID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUnit = `-- name: CreateUnit :one
INSERT INTO units (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateUnit(ctx context.Context, name string) (Unit, error) {
	row := q.db.QueryRow(ctx, createUnit, name)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, surname, email, password_hash, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, surname, email, password_hash, role, created_at, updated_at
`

type CreateUserParams struct {
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         UserRoles `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Surname,
		arg.Email,
		arg.PasswordHash,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserNotificationTg = `-- name: CreateUserNotificationTg :one
INSERT INTO user_notification_tgs (id, tg_username)
VALUES ($1, $2)
RETURNING id, tg_username, created_at, updated_at
`

type CreateUserNotificationTgParams struct {
	ID         int64  `json:"id"`
	TgUsername string `json:"tg_username"`
}

func (q *Queries) CreateUserNotificationTg(ctx context.Context, arg CreateUserNotificationTgParams) (UserNotificationTg, error) {
	row := q.db.QueryRow(ctx, createUserNotificationTg, arg.ID, arg.TgUsername)
	var i UserNotificationTg
	err := row.Scan(
		&i.ID,
		&i.TgUsername,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserTg = `-- name: CreateUserTg :one
INSERT INTO user_tgs (id, name, tg_username)
VALUES ($1, $2, $3)
RETURNING id, name, tg_username, unit_id, created_at, updated_at
`

type CreateUserTgParams struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	TgUsername string `json:"tg_username"`
}

func (q *Queries) CreateUserTg(ctx context.Context, arg CreateUserTgParams) (UserTg, error) {
	row := q.db.QueryRow(ctx, createUserTg, arg.ID, arg.Name, arg.TgUsername)
	var i UserTg
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TgUsername,
		&i.UnitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteApplication = `-- name: DeleteApplication :exec
DELETE FROM applications WHERE id = $1
`

func (q *Queries) DeleteApplication(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteApplication, id)
	return err
}

const deleteUnit = `-- name: DeleteUnit :exec
DELETE FROM units WHERE id = $1
`

func (q *Queries) DeleteUnit(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUnit, id)
	return err
}

const deleteUserNotificationsTg = `-- name: DeleteUserNotificationsTg :exec
DELETE FROM user_notification_tgs WHERE id = $1
`

func (q *Queries) DeleteUserNotificationsTg(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUserNotificationsTg, id)
	return err
}

const getApplications = `-- name: GetApplications :many
SELECT 
    applications.id,
    applications.text,
    applications.status,
    user_tgs.id, user_tgs.name, user_tgs.tg_username, user_tgs.unit_id, user_tgs.created_at, user_tgs.updated_at,
    units.id, units.name, units.created_at, units.updated_at
FROM applications
LEFT JOIN user_tgs ON applications.user_tg_id = user_tgs.id
LEFT JOIN units ON applications.unit_id = units.id
ORDER BY applications.created_at DESC
`

type GetApplicationsRow struct {
	ID     uuid.UUID         `json:"id"`
	Text   string            `json:"text"`
	Status ApplicationStatus `json:"status"`
	UserTg UserTg            `json:"user_tg"`
	Unit   Unit              `json:"unit"`
}

func (q *Queries) GetApplications(ctx context.Context) ([]GetApplicationsRow, error) {
	rows, err := q.db.Query(ctx, getApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetApplicationsRow
	for rows.Next() {
		var i GetApplicationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.Status,
			&i.UserTg.ID,
			&i.UserTg.Name,
			&i.UserTg.TgUsername,
			&i.UserTg.UnitID,
			&i.UserTg.CreatedAt,
			&i.UserTg.UpdatedAt,
			&i.Unit.ID,
			&i.Unit.Name,
			&i.Unit.CreatedAt,
			&i.Unit.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getApplicationsByUnitID = `-- name: GetApplicationsByUnitID :many
SELECT id, text, status, unit_id, user_tg_id, created_at, updated_at FROM applications WHERE unit_id = $1 ORDER BY created_at DESC
`

func (q *Queries) GetApplicationsByUnitID(ctx context.Context, unitID uuid.UUID) ([]Application, error) {
	rows, err := q.db.Query(ctx, getApplicationsByUnitID, unitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Application
	for rows.Next() {
		var i Application
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.Status,
			&i.UnitID,
			&i.UserTgID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnits = `-- name: GetUnits :many
SELECT id, name, created_at, updated_at FROM units ORDER BY created_at DESC
`

func (q *Queries) GetUnits(ctx context.Context) ([]Unit, error) {
	rows, err := q.db.Query(ctx, getUnits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Unit
	for rows.Next() {
		var i Unit
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, surname, email, password_hash, role, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, surname, email, password_hash, role, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserNotificationsTg = `-- name: GetUserNotificationsTg :one
SELECT id, tg_username, created_at, updated_at FROM user_notification_tgs WHERE id = $1
`

func (q *Queries) GetUserNotificationsTg(ctx context.Context, id int64) (UserNotificationTg, error) {
	row := q.db.QueryRow(ctx, getUserNotificationsTg, id)
	var i UserNotificationTg
	err := row.Scan(
		&i.ID,
		&i.TgUsername,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserNotificationsTgIds = `-- name: GetUserNotificationsTgIds :many
SELECT id FROM user_notification_tgs
`

func (q *Queries) GetUserNotificationsTgIds(ctx context.Context) ([]int64, error) {
	rows, err := q.db.Query(ctx, getUserNotificationsTgIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTg = `-- name: GetUserTg :one
SELECT u.id, u.name AS user_name, u.unit_id, un.name AS unit_name
FROM user_tgs u
LEFT JOIN units un ON u.unit_id = un.id
WHERE u.id = $1
`

type GetUserTgRow struct {
	ID       int64       `json:"id"`
	UserName string      `json:"user_name"`
	UnitID   uuid.UUID   `json:"unit_id"`
	UnitName pgtype.Text `json:"unit_name"`
}

func (q *Queries) GetUserTg(ctx context.Context, id int64) (GetUserTgRow, error) {
	row := q.db.QueryRow(ctx, getUserTg, id)
	var i GetUserTgRow
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.UnitID,
		&i.UnitName,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, surname, email, password_hash, role, created_at, updated_at FROM users ORDER BY created_at DESC
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Surname,
			&i.Email,
			&i.PasswordHash,
			&i.Role,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateApplicationStatus = `-- name: UpdateApplicationStatus :one
UPDATE applications SET status = $1 WHERE id = $2 RETURNING id, text, status, unit_id, user_tg_id, created_at, updated_at
`

type UpdateApplicationStatusParams struct {
	Status ApplicationStatus `json:"status"`
	ID     uuid.UUID         `json:"id"`
}

func (q *Queries) UpdateApplicationStatus(ctx context.Context, arg UpdateApplicationStatusParams) (Application, error) {
	row := q.db.QueryRow(ctx, updateApplicationStatus, arg.Status, arg.ID)
	var i Application
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.Status,
		&i.UnitID,
		&i.UserTgID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUnitName = `-- name: UpdateUnitName :one
UPDATE units SET name = $1 WHERE id = $2 RETURNING id, name, created_at, updated_at
`

type UpdateUnitNameParams struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUnitName(ctx context.Context, arg UpdateUnitNameParams) (Unit, error) {
	row := q.db.QueryRow(ctx, updateUnitName, arg.Name, arg.ID)
	var i Unit
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users SET role = $1 WHERE id = $2 RETURNING id, name, surname, email, password_hash, role, created_at, updated_at
`

type UpdateUserRoleParams struct {
	Role UserRoles `json:"role"`
	ID   uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserRole, arg.Role, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserTgName = `-- name: UpdateUserTgName :exec
UPDATE user_tgs
SET name = $1
WHERE id = $2
`

type UpdateUserTgNameParams struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateUserTgName(ctx context.Context, arg UpdateUserTgNameParams) error {
	_, err := q.db.Exec(ctx, updateUserTgName, arg.Name, arg.ID)
	return err
}

const updateUserUnitID = `-- name: UpdateUserUnitID :exec
UPDATE user_tgs SET unit_id = $1 WHERE id = $2
`

type UpdateUserUnitIDParams struct {
	UnitID uuid.UUID `json:"unit_id"`
	UserID int64     `json:"user_id"`
}

func (q *Queries) UpdateUserUnitID(ctx context.Context, arg UpdateUserUnitIDParams) error {
	_, err := q.db.Exec(ctx, updateUserUnitID, arg.UnitID, arg.UserID)
	return err
}
