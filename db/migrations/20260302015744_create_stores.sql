-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS stores (
  id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name         TEXT NOT NULL,
  created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at   TIMESTAMPTZ,
  deleted_at   TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_stores_deleted_at ON stores(deleted_at);

-- +goose Down
DROP TABLE IF EXISTS stores;