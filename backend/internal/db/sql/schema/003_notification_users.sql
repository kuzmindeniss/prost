-- +goose Up
CREATE TABLE user_notification_tgs (
  id BIGINT PRIMARY KEY,
  tg_username text NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON user_notification_tgs
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- +goose Down
DROP TABLE IF EXISTS user_notification_tgs;
DROP TRIGGER IF EXISTS trigger_update_updated_at ON user_notification_tgs;
