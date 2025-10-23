-- name: UpdateUserPassword :one
UPDATE users
SET senha = $2, updated_at = CURRENT_TIMESTAMP
WHERE email = $1
RETURNING id, nome, email, senha, created_at, updated_at;

