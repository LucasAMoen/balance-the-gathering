-- name: GetCards :many
SELECT * FROM cards;

-- name: GetCardById :one
SELECT * FROM cards
WHERE id = $1;