DROP TABLE services;

ALTER TABLE employees DROP CONSTRAINT employees_garage_id_fkey;

DROP TABLE garages;

DROP TABLE employees;

DROP TYPE roles;
