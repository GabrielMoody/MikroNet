CREATE TYPE roles AS ENUM ('user', 'driver', 'admin', 'government', 'business_owner');
CREATE TYPE genders AS ENUM ('male', 'female');
CREATE TYPE statuses AS ENUM ('on', 'off');

CREATE TABLE IF NOT EXISTS Users (
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
     oauth VARCHAR(255) DEFAULT NULL,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE BusinessOwners (
        id uuid PRIMARY KEY,
        NIK VARCHAR(255),
        verified BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Drivers (
           id uuid PRIMARY KEY,
           owner_id uuid,
           registration_number VARCHAR(255) UNIQUE,
           status statuses DEFAULT 'off',
           latitude DECIMAL(10, 8),
           longitude DECIMAL(11, 8),
           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           FOREIGN KEY (id) REFERENCES Users(id),
           FOREIGN KEY (owner_id) REFERENCES BusinessOwners(id)
);

CREATE TABLE IF NOT EXISTS Trips (
         id uuid PRIMARY KEY,
         user_id uuid,
         driver_id uuid,
         location VARCHAR(255),
         destination VARCHAR(255),
         trip_date TIMESTAMP,
         created_at TIMESTAMP,
         FOREIGN KEY (user_id) REFERENCES Users(id),
         FOREIGN KEY (driver_id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS Reviews (
       id uuid PRIMARY KEY ,
       user_id uuid,
       driver_id uuid,
       review VARCHAR(255),
       star INT,
       created_at TIMESTAMP DEFAULT NOW(),
       FOREIGN KEY (user_id) REFERENCES Users(id),
       FOREIGN KEY (driver_id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS Routes (
      id uuid PRIMARY KEY,
      route_name VARCHAR(255),
      initial_route VARCHAR(255),
      destination_route VARCHAR(255),
      created_at TIMESTAMP,
      FOREIGN KEY (id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS ResetPassword (
     id INT PRIMARY KEY,
     user_id uuid,
     reset_code VARCHAR(255),
     created_at TIMESTAMP
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
