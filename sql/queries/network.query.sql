-- name: GetNetwork :one
SELECT * FROM networks WHERE id = $1 LIMIT 1;

-- name: GetNetworks :many
SELECT
  *
from networks
  LIMIT sqlc.narg('limit')
  OFFSET sqlc.narg('offset');

-- name: CreateNetwork :one
INSERT INTO networks (code, name) VALUES ($1, $2) RETURNING *;

-- name: UpdateNetwork :one
UPDATE networks SET code = $1, name = $2 WHERE id = $3 RETURNING *;

-- name: DeleteNetwork :one
DELETE FROM networks WHERE id = $1 RETURNING *;
