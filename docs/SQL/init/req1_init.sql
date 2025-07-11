DROP DATABASE IF EXISTS frescos_db;
CREATE DATABASE frescos_db;
USE frescos_db;
SET SQL_SAFE_UPDATES = 0;

DROP TABLE IF EXISTS localities;
CREATE TABLE localities(
    id INT PRIMARY KEY AUTO_INCREMENT,
    locality_name VARCHAR(255),
    province_name VARCHAR(255),
    country_name VARCHAR(255)
);

DROP TABLE IF EXISTS sellers;
CREATE TABLE sellers(
    id INT PRIMARY KEY AUTO_INCREMENT,
    cid VARCHAR(255) NOT NULL CHECK (cid > 0),
    company_name VARCHAR(255) NOT NULL,
    `address` VARCHAR(255)  NOT NULL,
    telephone VARCHAR(255)  NOT NULL,
    locality_id INT  NOT NULL,

    FOREIGN KEY(locality_id) REFERENCES localities(id),

    UNIQUE(cid)
);
