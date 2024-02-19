-- name: CreateUser :one
INSERT INTO users (
    name,
    last_name,
    username,
    password,
    email
)
VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
where email = $1;

-- name: GetUserById :one
SELECT * FROM users
where id = $1;

-- name: VerifyUser :execrows
UPDATE users
SET verified = true
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = $1,
    last_name = $2, 
    username = $3,
    password =$4,
    email = $5,
    plan_id = $6,
    last_updated = now()
WHERE id = $7;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;


-- name: GetUserPlanById :one
SELECT * FROM users
INNER JOIN plans ON users.plan_id = plans.id
WHERE users.id = $1;