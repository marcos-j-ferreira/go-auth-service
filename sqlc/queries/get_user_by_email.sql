-- name: GetUserByEmail :one
SELECT id, nome, email, senha, created_at, updated_at
FROM users
WHERE email = $1;

