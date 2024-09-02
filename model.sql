SET GLOBAL event_scheduler = on;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_of_birth TIME,
    age INTEGER,
    role VARCHAR(255) DEFAULT 'user' CHECK(role in ('user', 'driver', 'admin')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS driver_detail (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255),
    registration_number VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS trips (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255),
    driver_id VARCHAR(255),
    location VARCHAR(255),
    destination VARCHAR(255),
    trip_date TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id VARCHAR(255) PRIMARY KEY ,
    user_id VARCHAR(255),
    driver_id VARCHAR(255),
    review VARCHAR(255),
    star INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS routes (
    id VARCHAR(255) PRIMARY KEY,
    route_name VARCHAR(255),
    initial_route VARCHAR(255),
    destination_route VARCHAR(255),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reset_password (
    id INT PRIMARY KEY,
    user_id VARCHAR(255),
    reset_code VARCHAR(255),
    created_at TIMESTAMP
);

CREATE EVENT delete_reset_password_link
ON SCHEDULE EVERY 1 HOUR
ON COMPLETION PRESERVE

DO BEGIN
    DELETE FROM reset_password WHERE created_at < NOW() - INTERVAL 1 HOUR;
end;