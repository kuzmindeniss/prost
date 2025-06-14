-- +goose Up
CREATE TABLE units (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE user_tgs (
  id BIGINT PRIMARY KEY,
  name text NOT NULL,
  tg_username text NOT NULL,
  unit_id UUID REFERENCES units(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TYPE application_status AS ENUM ('pending', 'done');

CREATE TABLE applications (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  text text NOT NULL,
  status application_status NOT NULL DEFAULT 'pending',
  unit_id UUID REFERENCES units(id) ON DELETE RESTRICT,
  user_tg_id BIGINT REFERENCES user_tgs(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TYPE user_roles AS ENUM ('admin', 'user');

CREATE TABLE users (
  id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  surname text NOT NULL,
  email text NOT NULL UNIQUE,
  password_hash text NOT NULL,
  role  USER_ROLES NOT NULL DEFAULT 'user',
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

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON units
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON user_tgs
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON applications
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- +goose Down
DROP TABLE IF EXISTS applications;
DROP TABLE IF EXISTS user_tgs;
DROP TABLE IF EXISTS units;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON users;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON units;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON user_tgs;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON applications;
DROP FUNCTION IF EXISTS update_updated_at;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_roles;
DROP TYPE IF EXISTS application_status;
