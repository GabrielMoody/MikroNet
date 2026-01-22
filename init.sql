CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS authentications (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    username VARCHAR(255) UNIQUE,
    fullname VARCHAR(255),
    phone_number VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS drivers (
    id BIGINT PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    name VARCHAR(255),
    phone_number VARCHAR(255),
    vehicle_type VARCHAR(100),
    plate_number VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS driver_status (
    driver_id INT PRIMARY KEY REFERENCES drivers(id),
    is_online BOOLEAN default false,
    is_busy BOOLEAN default false,
    last_activity_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS driver_locations (
    driver_id INT PRIMARY KEY REFERENCES drivers(id),
    location GEOGRAPHY(POINT, 4326),
    heading FLOAT,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    driver_id INT REFERENCES drivers(id),
    pickup_point GEOMETRY(POINT, 4326),
    dropoff_point GEOMETRY(POINT, 4326),
    status VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_driver_locations_geom
ON driver_locations
USING GIST (location);
