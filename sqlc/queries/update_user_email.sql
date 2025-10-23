-- name: UpdateUserEmail :one
UPDATE users
SET email = $2, update_at = CURRENT_TIMESTAMP
WHERE email = $1 
RETURNING id, nome, email, senha, created_at, updated_at;

