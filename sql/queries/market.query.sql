-- name: GetMarket :one
SELECT * FROM markets WHERE id = $1 LIMIT 1;

-- name: GetMarkets :many
SELECT
  *
from markets
  LIMIT sqlc.narg('limit')
  OFFSET sqlc.narg('offset');

-- name: CreateMarket :one
INSERT INTO markets (name, description) VALUES ($1, $2) RETURNING *;

-- name: UpdateMarket :one
UPDATE markets SET name = $1, description = $2 WHERE id = $3 RETURNING *;

-- name: DeleteMarket :one
DELETE FROM markets WHERE id = $1 RETURNING *;
