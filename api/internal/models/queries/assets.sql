-- name: CreateAsset :one
INSERT INTO assets(model_id, "serial", comments)
VALUES (@model_id, @serial, @comments)
RETURNING *;

-- name: DeleteAssetByID :execrows
DELETE FROM assets
WHERE id = @asset_id;

-- name: GetAssetByID :one
SELECT * FROM assets WHERE id = @asset_id;

-- name: ListAssets :many
SELECT * FROM assets ORDER BY id LIMIT 50;

-- name: UpdateAssetByID :one
UPDATE assets
SET model_id = @model_id, "serial" = @serial, comments = @comments
WHERE id = @asset_id
RETURNING *;
