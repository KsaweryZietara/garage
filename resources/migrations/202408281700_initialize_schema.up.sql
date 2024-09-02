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
    garage_id INT
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
    garage_id INT REFERENCES garages(id)
);
