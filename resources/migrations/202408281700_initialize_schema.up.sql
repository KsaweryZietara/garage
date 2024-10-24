DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'roles') THEN
        CREATE TYPE roles AS ENUM ('OWNER', 'MECHANIC');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS employees
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    role roles NOT NULL,
    garage_id INT,
    confirmed BOOLEAN,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS garages
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    number VARCHAR(15) NOT NULL,
    postal_code VARCHAR(15) NOT NULL,
    phone_number VARCHAR(15) unique NOT NULL,
    latitude DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL,
    logo BYTEA,
    owner_id INT REFERENCES employees(id)
);

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'employees_garage_id_fkey'
    ) THEN
        ALTER TABLE employees ADD CONSTRAINT employees_garage_id_fkey
            FOREIGN KEY (garage_id) REFERENCES garages(id);
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS services
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    time int NOT NULL,
    price int NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    garage_id INT REFERENCES garages(id)
);

CREATE TABLE IF NOT EXISTS confirmation_codes
(
    id UUID PRIMARY KEY,
    employee_id INT REFERENCES employees(id)
);

CREATE TABLE IF NOT EXISTS customers
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS makes
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS models
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    make_id INT REFERENCES makes(id)
);

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    service_id INT REFERENCES services(id),
    employee_id INT REFERENCES employees(id),
    customer_id INT REFERENCES customers(id),
    model_id INT REFERENCES models(id)
);
