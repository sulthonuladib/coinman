-- name: GetCoin :one
SELECT * FROM coins WHERE id = $1 LIMIT 1;

-- name: GetCoins :many
SELECT
  *
  FROM coins
  WHERE
  symbol ILIKE '%' || sqlc.narg('symbol') || '%' OR name ILIKE '%' || sqlc.narg('symbol') || '%'
  LIMIT sqlc.narg('limit')
  OFFSET sqlc.narg('offset');

-- name: CreateCoin :one
INSERT INTO coins (symbol, name) VALUES ($1, $2) RETURNING *;

-- name: UpdateCoin :one
UPDATE coins SET symbol = $1, name = $2 WHERE id = $3 RETURNING *;

-- name: DeleteCoin :one
DELETE FROM coins WHERE id = $1 RETURNING *;

-- name: SearchCoins :many
SELECT * FROM coins WHERE symbol ILIKE '%' || sqlc.narg('symbol') || '%' OR name ILIKE '%' || sqlc.narg('symbol') || '%' LIMIT sqlc.narg('limit') OFFSET sqlc.narg('offset');
