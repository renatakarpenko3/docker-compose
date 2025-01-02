CREATE DATABASE testdb;
USE testdb;
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255) 
);

INSERT INTO users (id, name) VALUES (1, 'Bob');
INSERT INTO users (id, name) VALUES (2, 'Bella');
USE testdb;
