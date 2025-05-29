-- +goose Up
CREATE TABLE units (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL
);

CREATE TABLE users_tg (
  id BIGINT PRIMARY KEY,
  name text NOT NULL,
  tg_username text NOT NULL,
  unit_id UUID REFERENCES units(id) ON DELETE SET NULL
);

CREATE TYPE application_status AS ENUM ('pending', 'done');

CREATE TABLE applications (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  text text NOT NULL,
  status application_status NOT NULL DEFAULT 'pending',
  unit_id UUID REFERENCES units(id) ON DELETE SET NULL,
  user_tg_id BIGINT REFERENCES users_tg(id) ON DELETE SET NULL
);

CREATE TYPE user_roles AS ENUM ('admin', 'user');

CREATE TABLE users (
  id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  surname text NOT NULL,
  email text NOT NULL UNIQUE,
  password_hash text NOT NULL,
  role  USER_ROLES DEFAULT 'user',
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- +goose Down
DROP TABLE IF EXISTS applications;
DROP TABLE IF EXISTS users_tg;
DROP TABLE IF EXISTS units;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_roles;
