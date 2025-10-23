-- name: CreateUser :one
INSERT INTO users (
    nome,
    email,
    senha
) VALUES (
    $1, $2, $3
)
RETURNING id, nome, email, created_at, updated_at;