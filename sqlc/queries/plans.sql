-- name: GetPlanById :one
SELECT * FROM plans
where id = $1;

-- name: GetPlanByName :one
SELECT * FROM plans
where name = $1;

-- name: GetAllPlans :many
SELECT * FROM plans;