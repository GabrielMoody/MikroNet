CREATE TABLE IF NOT EXISTS authentications (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    username VARCHAR(255) UNIQUE,
    fullname VARCHAR(255),
    phone_number VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS drivers (
    id BIGSERIAL PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    name VARCHAR(255),
    phone_number VARCHAR(255),
    vehicle_type VARCHAR(100),
    plate_number VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS driver_status (
    driver_id BIGSERIAL PRIMARY KEY REFERENCES drivers(id),
    is_online BOOLEAN default false,
    is_busy BOOLEAN default false,
    last_activity_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS driver_locations (
    driver_id BIGSERIAL PRIMARY KEY REFERENCES drivers(id),
    location TEXT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    driver_id INTEGER REFERENCES drivers(id),
    pickup_point TEXT,
    dropoff_point TEXT,
    status VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);