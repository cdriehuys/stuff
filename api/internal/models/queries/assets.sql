-- name: CreateAsset :one
WITH model AS (SELECT * FROM models WHERE id = @model_id)
INSERT INTO assets(model_id, "serial", comments)
VALUES (@model_id, @serial, @comments)
RETURNING sqlc.embed(assets), (SELECT vendor_id FROM model) AS vendor_id;

-- name: DeleteAssetByID :execrows
DELETE FROM assets
WHERE id = @asset_id;

-- name: GetAssetByID :one
SELECT sqlc.embed(a), m.vendor_id
FROM assets a
    JOIN models m ON a.model_id = m.id
WHERE a.id = @asset_id;

-- name: ListAssets :many
SELECT sqlc.embed(a), m.vendor_id
FROM assets a
    JOIN models m ON a.model_id = m.id
ORDER BY a.id
LIMIT 50;

-- name: ListAssetsByModel :many
SELECT sqlc.embed(a), m.vendor_id
FROM assets a
    JOIN models m ON a.model_id = m.id
WHERE a.model_id = @model_id
ORDER BY a.id
LIMIT 50;

-- name: UpdateAssetByID :one
WITH model AS (SELECT * FROM models WHERE id = @model_id)
UPDATE assets a
SET model_id = @model_id, "serial" = @serial, comments = @comments
WHERE a.id = @asset_id
RETURNING sqlc.embed(a), (SELECT vendor_id FROM model) AS vendor_id;
