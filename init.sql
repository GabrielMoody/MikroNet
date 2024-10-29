CREATE TYPE roles AS ENUM ('user', 'driver', 'admin', 'government', 'business_owner');
CREATE TYPE genders AS ENUM ('male', 'female', '');
CREATE TYPE statuses AS ENUM ('on', 'off');
CREATE TYPE order_status AS ENUM  ('pending', 'accepted', 'completed', 'canceled');

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    age INTEGER,
    gender genders,
    role roles NOT NULL,
    is_blocked BOOLEAN DEFAULT FALSE,
    image_url VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE business_owners (
    id uuid PRIMARY KEY,
    NIK VARCHAR(255),
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS routes (
    id uuid PRIMARY KEY,
    route_name VARCHAR(255),
    initial_route VARCHAR(255),
    destination_route VARCHAR(255),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS drivers (
    id uuid PRIMARY KEY,
    owner_id uuid,
    route_id uuid,
    registration_number VARCHAR(255) UNIQUE,
    status statuses DEFAULT 'off',
    available_seats INT CONSTRAINT seat_constraint CHECK ( available_seats <= 9 ) DEFAULT 9,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES users(id),
    FOREIGN KEY (owner_id) REFERENCES business_owners(id),
    FOREIGN KEY (route_id) REFERENCES routes(id)
);

CREATE TABLE IF NOT EXISTS driver_location (
    driver_id uuid PRIMARY KEY,
    location geography(Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS orders
(
    id uuid PRIMARY KEY,
    user_id uuid,
    driver_id uuid,
    start_location geography(Point, 4326),
    end_location geography(Point, 4326),
    status order_status,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS driver_location_logs (
    id uuid PRIMARY KEY,
    location geography(Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS passenger_histories (
    id uuid PRIMARY KEY,
    start_location geography(Point, 4326),
    end_location geography(Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id uuid PRIMARY KEY,
    review VARCHAR(255),
    star INT CONSTRAINT star_constraint CHECK (star BETWEEN 1 AND 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS reset_password (
    id BIGSERIAL PRIMARY KEY,
    user_id uuid,
    reset_code VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id uuid,
    title VARCHAR(255),
    message VARCHAR(225),
    is_read BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE FUNCTION expire_reset_password_links() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
DELETE FROM ResetPassword WHERE current_timestamp < NOW() - INTERVAL '1 minute';
RETURN NEW;
END;
$$;

CREATE TRIGGER expire_reset_password_links
    AFTER INSERT ON ResetPassword
    EXECUTE PROCEDURE expire_reset_password_links();

CREATE OR REPLACE FUNCTION log_driver_location_table_changes()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO driver_location_logs(id, location)
    VALUES (OLD.driver_id, OLD.location);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_driver_location_table_updates
    BEFORE UPDATE ON driver_location
    FOR EACH ROW
EXECUTE FUNCTION log_driver_location_table_changes();