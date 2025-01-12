-- name: CreateModel :one
INSERT INTO models(model, vendor_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteModelByID :execrows
DELETE FROM models
WHERE id = $1;

-- name: GetModelByID :one
SELECT * FROM models WHERE id = $1;

-- name: ListModels :many
SELECT * FROM models ORDER BY model LIMIT 50;

-- name: ListModelsByVendorID :many
SELECT * FROM models
WHERE vendor_id = $1
ORDER BY model LIMIT 50;

-- name: UpdateModelByID :one
UPDATE models
SET vendor_id = @vendor_id, model = $2, name = $3
WHERE id = $1
RETURNING *;
