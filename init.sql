CREATE TYPE roles AS ENUM ('user', 'driver', 'admin', 'government', 'business_owner');
CREATE TYPE genders AS ENUM ('male', 'female', '');
CREATE TYPE statuses AS ENUM ('on', 'off');
CREATE TYPE order_status AS ENUM  ('pending', 'accepted', 'completed', 'canceled');

CREATE TABLE IF NOT EXISTS auth (
    id uuid PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS routes (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS driver_details (
    id uuid PRIMARY KEY,
    route_id uuid,
    name VARCHAR(255),
    phone_number VARCHAR(255) UNIQUE,
    sim VARCHAR(255) UNIQUE,
    license_number VARCHAR(255) UNIQUE,
    status statuses DEFAULT 'off',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES users(id),
    FOREIGN KEY (route_id) REFERENCES routes(id)
);

CREATE TABLE IF NOT EXISTS driver_location (
    driver_id uuid PRIMARY KEY,
    location geography(Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS driver_location_logs (
    id uuid PRIMARY KEY,
    location geography(Point, 4326),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES drivers(id)
);

CREATE TABLE IF NOT EXISTS passenger_details (
    id uuid PRIMARY KEY,
    name VARCHAR(255),
    FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id uuid PRIMARY KEY,
    passenger_id uuid,
    driver_id uuid,
    review VARCHAR(255),
    star INT CONSTRAINT star_constraint CHECK (star BETWEEN 1 AND 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (passenger_id) REFERENCES passenger_details(id),
    FOREIGN KEY (driver_id) REFERENCES driver_details(id)
);

CREATE TABLE IF NOT EXISTS reset_password (
    id INT PRIMARY KEY,
    user_id uuid,
    reset_code VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE blocked_account {
    id INT PRIMARY KEY,
    account_id uuid,
    role VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES users(id)
}

CREATE TABLE IF NOT EXISTS notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id uuid,
    title VARCHAR(255),
    message VARCHAR(225),
    is_read BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DELIMITER $$

CREATE TRIGGER expire_reset_password_links
    AFTER INSERT ON reset_passwords
    FOR EACH ROW
BEGIN
    DELETE FROM reset_passwords WHERE created_at < NOW() - INTERVAL 1 HOUR;
END$$

DELIMITER ;

CREATE OR REPLACE FUNCTION log_driver_location_table_changes()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO driver_location_logs(driver_id, location)
    VALUES (OLD.driver_id, OLD.location);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_driver_location_table_updates
    BEFORE UPDATE ON driver_locations
    FOR EACH ROW
EXECUTE FUNCTION log_driver_location_table_changes();

CREATE EXTENSION postgis;
CREATE EXTENSION postgis_topology;
INSERT INTO routes ( id, route_name) values
    (1, 'terminal malalayang- pusat kota'),
    (2, 'terminal malalayang - terminal karombasan'),
    (3, 'terminal karombasan- pusat kota'),
    (4, 'terminal karombasan- pusat kota'),
    (5, 'winangun- pusat kota'),
    (6, 'karombasan - pusat kota'),
    (7, 'Terminal karombasan - terminal paal dua'),
    (8, 'terminal paal dua - pusat kota'),
    (9, 'kairagi- pusat kota'),
    (10, 'perkamil - pusat kota'),
    (11, 'banjer, paal 4, taas- pusat kota'),
    (12, 'wonasa - pusat kota'),
    (13, 'tuminting - pusat kota'),
    (14, 'terminal paal dua - lapangan'),
    (15, 'terminal paal dua- politeknik)'),
    (16, 'tuminting- pandu'),
    (17, 'tuminting- tongkaina')

DELIMITER //

CREATE TRIGGER set_username_from_email
BEFORE INSERT ON passenger_details
FOR EACH ROW
BEGIN
  DECLARE email_val VARCHAR(255);

  -- Get email from user_accounts using the user_id
  SELECT email INTO email_val
  FROM users
  WHERE users = NEW.id;

  -- Set username only if not manually provided
  IF NEW.name IS NULL OR NEW.name = '' THEN
    SET NEW.name = SUBSTRING_INDEX(email_val, '@', 1);
  END IF;
END;
//

DELIMITER ;
