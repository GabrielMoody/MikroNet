CREATE TABLE IF NOT EXISTS authentications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(100),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    username VARCHAR(255) UNIQUE,
    fullname VARCHAR(255),
    phone_number VARCHAR(100),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS drivers (
    id INTEGER PRIMARY KEY REFERENCES authentications(id) ON DELETE CASCADE,
    name VARCHAR(255),
    phone_number VARCHAR(255),
    vehicle_type VARCHAR(100),
    plate_number VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS driver_status (
    driver_id INTEGER PRIMARY KEY REFERENCES drivers(id),
    is_online BOOLEAN default false,
    is_busy BOOLEAN default false,
    last_activity_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS driver_locations (
    driver_id INTEGER PRIMARY KEY REFERENCES drivers(id),
    location TEXT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER REFERENCES users(id),
    driver_id INTEGER REFERENCES drivers(id),
    pickup_point TEXT,
    dropoff_point TEXT,
    status VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);