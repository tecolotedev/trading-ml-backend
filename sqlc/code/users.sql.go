// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    last_name,
    username,
    password,
    email
)
VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, last_name, username, password, email, verified, created_at, last_updated, plan_id
`

type CreateUserParams struct {
	Name     string
	LastName string
	Username string
	Password string
	Email    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.LastName,
		arg.Username,
		arg.Password,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Verified,
		&i.CreatedAt,
		&i.LastUpdated,
		&i.PlanID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, last_name, username, password, email, verified, created_at, last_updated, plan_id FROM users
where email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Verified,
		&i.CreatedAt,
		&i.LastUpdated,
		&i.PlanID,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, last_name, username, password, email, verified, created_at, last_updated, plan_id FROM users
where id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Verified,
		&i.CreatedAt,
		&i.LastUpdated,
		&i.PlanID,
	)
	return i, err
}

const getUserPlanById = `-- name: GetUserPlanById :one
SELECT users.id, users.name, last_name, username, password, email, verified, users.created_at, users.last_updated, plan_id, plans.id, plans.name, max_historical_bars, max_symbols, max_indicators_per_symbol, max_models, plans.created_at, plans.last_updated FROM users
INNER JOIN plans ON users.plan_id = plans.id
WHERE users.id = $1
`

type GetUserPlanByIdRow struct {
	ID                     int32
	Name                   string
	LastName               string
	Username               string
	Password               string
	Email                  string
	Verified               pgtype.Bool
	CreatedAt              pgtype.Timestamp
	LastUpdated            pgtype.Timestamp
	PlanID                 pgtype.Int4
	ID_2                   int32
	Name_2                 pgtype.Text
	MaxHistoricalBars      pgtype.Int4
	MaxSymbols             pgtype.Int4
	MaxIndicatorsPerSymbol pgtype.Int4
	MaxModels              pgtype.Int4
	CreatedAt_2            pgtype.Timestamp
	LastUpdated_2          pgtype.Timestamp
}

func (q *Queries) GetUserPlanById(ctx context.Context, id int32) (GetUserPlanByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserPlanById, id)
	var i GetUserPlanByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Verified,
		&i.CreatedAt,
		&i.LastUpdated,
		&i.PlanID,
		&i.ID_2,
		&i.Name_2,
		&i.MaxHistoricalBars,
		&i.MaxSymbols,
		&i.MaxIndicatorsPerSymbol,
		&i.MaxModels,
		&i.CreatedAt_2,
		&i.LastUpdated_2,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $1,
    last_name = $2, 
    username = $3,
    password =$4,
    email = $5,
    last_updated = now()
WHERE id = $6
`

type UpdateUserParams struct {
	Name     string
	LastName string
	Username string
	Password string
	Email    string
	ID       int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Name,
		arg.LastName,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.ID,
	)
	return err
}

const updateUserPlan = `-- name: UpdateUserPlan :exec
UPDATE users
SET 
    plan_id = $1,
    last_updated = now()
WHERE id = $2
`

type UpdateUserPlanParams struct {
	PlanID pgtype.Int4
	ID     int32
}

func (q *Queries) UpdateUserPlan(ctx context.Context, arg UpdateUserPlanParams) error {
	_, err := q.db.Exec(ctx, updateUserPlan, arg.PlanID, arg.ID)
	return err
}

const verifyUser = `-- name: VerifyUser :execrows
UPDATE users
SET verified = true
WHERE id = $1
`

func (q *Queries) VerifyUser(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, verifyUser, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
