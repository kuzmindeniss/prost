-- +goose Up
-- Create default admin user
-- Password hash for 'ah23khsl343t_23jks' generated with bcrypt
INSERT INTO users (name, surname, email, password_hash, role)
VALUES ('Admin', 'User', 'admin@admin.com', '$2a$10$nmGqDuP1Dv7AWxg5V/.4beAB4LzqYtXS2NM8BV8LQSJ/.7VfqAF/a', 'admin')
ON CONFLICT (email) DO NOTHING;

-- +goose Down
-- Remove the default admin user (if needed)
DELETE FROM users WHERE email = 'admin@admin.com'; 
