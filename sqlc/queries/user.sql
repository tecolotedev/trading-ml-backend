-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
)
VALUES (
    $1, $2, $3
) RETURNING id,username,email,verified,created_at;

-- name: GetUser :one
SELECT * FROM users
where email = $1;

-- name: GetUserById :one
SELECT * FROM users
where id = $1;

-- name: VerifyUser :execrows
UPDATE users
SET verified = true
WHERE id = $1;