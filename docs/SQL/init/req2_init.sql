DROP DATABASE IF EXISTS warehouse_db;
CREATE DATABASE IF NOT EXISTS warehouse_db;
USE warehouse_db;
SET SQL_SAFE_UPDATES = 0;

DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS carries;
DROP TABLE IF EXISTS localities;

CREATE TABLE localities (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE carries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(50) NOT NULL UNIQUE,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(50) NOT NULL,
    locality_id INT NOT NULL,
    FOREIGN KEY (locality_id) REFERENCES localities(id)
);
CREATE TABLE warehouses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    warehouse_code VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    minimum_capacity INT NOT NULL,
    minimum_temperature FLOAT NOT NULL
);



