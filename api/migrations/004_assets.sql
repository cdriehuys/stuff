CREATE TABLE assets (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    model_id BIGINT REFERENCES models(id) ON DELETE RESTRICT,
    "serial" TEXT NOT NULL,
    comments TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

SELECT _manage_updated_at('assets');

---- create above / drop below ----

DROP TABLE assets;
