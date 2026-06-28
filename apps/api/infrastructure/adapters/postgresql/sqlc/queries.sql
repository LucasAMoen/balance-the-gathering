-- name: GetCards :many
SELECT * FROM cards;

-- name: GetCardById :one
SELECT * FROM cards
WHERE id = $1;

-- name: InsertCards :copyfrom
INSERT INTO
cards (id, name, url, layout, imageUrl, price)
VALUES ($1,$2,$3,$4,$5,$6);