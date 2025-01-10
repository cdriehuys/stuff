-- name: CreateVendor :one
INSERT INTO vendors(name)
VALUES ($1)
RETURNING *;

-- name: DeleteVendorByID :execrows
DELETE FROM vendors
WHERE id = $1;

-- name: GetVendorByID :one
SELECT *
FROM vendors
WHERE id = $1;

-- name: ListVendors :many
SELECT * FROM vendors ORDER BY id LIMIT 50;
