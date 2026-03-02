-- +goose Up
CREATE TABLE IF NOT EXISTS employees (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  store_id      UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,

  name          TEXT NOT NULL,
  email         TEXT NOT NULL,
  password_hash TEXT NOT NULL,

  role          TEXT NOT NULL,
  is_active     BOOLEAN NOT NULL DEFAULT TRUE,

  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ,
  deleted_at    TIMESTAMPTZ,

  CONSTRAINT employees_role_check CHECK (role IN ('waiter', 'kitchen', 'cashier', 'manager')),
  CONSTRAINT employees_email_store_unique UNIQUE (store_id, email)
);

CREATE INDEX IF NOT EXISTS idx_employees_store_id ON employees(store_id);
CREATE INDEX IF NOT EXISTS idx_employees_store_id_role ON employees(store_id, role);
CREATE INDEX IF NOT EXISTS idx_employees_deleted_at ON employees(deleted_at);

-- +goose Down
DROP TABLE IF EXISTS employees;