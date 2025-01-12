CREATE TABLE models (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    model TEXT NOT NULL,
    vendor_id BIGINT NOT NULL REFERENCES vendors(id) ON DELETE RESTRICT,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(model, vendor_id)
);

SELECT _manage_updated_at('models');

---- create above / drop below ----

DROP TABLE models;
