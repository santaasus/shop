CREATE ROLE santaasus WITH LOGIN PASSWORD 'youShouldChangeThisPassword';

CREATE DATABASE shop_db OWNER santaasus;

GRANT ALL PRIVILEGES ON DATABASE shop_db TO santaasus;

\connect shop_db;

CREATE USER santaasus WITH PASSWORD 'youShouldChangeThisPassword';

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    user_name varchar(30),
    email varchar(30),
    first_name varchar(30),
    last_name varchar(30),
    hash_password varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)

CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name varchar(30) NULL
)

CREATE TABLE IF NOT EXISTS order (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGSERIAL PRIMARY KEY,
    is_payed BOOLEAN DEFAULT NULL 
)

-- The Function for update field `updated_at`
CREATE OR REPLACE FUNCTION on_user_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- The trigger for function `updated_at` before update `users`
CREATE TRIGGER on_user_updated_at_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION on_user_updated_at();

INSERT INTO users (id, user_name, email, first_name, last_name, hash_password, created_at, updated_at)
VALUES (1, 'santaasus', 'test@yandex.ru', 'Vladimir', 'S', '$2a$10$ARGDNUz.xsfWAaS2KCG2T.h5N3d9NTf77i0Q5dp6FdpYXSJI08ijW', '2024-07-26 05:23:20','2024-07-26 05:23:20');

