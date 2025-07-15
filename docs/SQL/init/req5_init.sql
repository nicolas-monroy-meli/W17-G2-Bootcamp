DROP DATABASE IF EXISTS frescos_db;
CREATE DATABASE frescos_db;
USE frescos_db;
SET SQL_SAFE_UPDATES = 0;

DROP TABLE IF EXISTS employees;
CREATE TABLE employees(
                           id INT PRIMARY KEY AUTO_INCREMENT,
                           id_card_number VARCHAR(255),
                           first_name VARCHAR(255),
                           last_name VARCHAR(255),
                            wareHouse_id INT NOT NULL,

                            --FOREIGN KEY(wareHouse_id) REFERENCES warehouse(id),
                            UNIQUE(id_card_number)
);

DROP TABLE IF EXISTS inbound_orders;
CREATE TABLE inbound_orders(
                        id INT PRIMARY KEY AUTO_INCREMENT,
                        order_date DATE,
                        order_number VARCHAR(255) NOT NULL,
                        employe_id INT  NOT NULL,
                        product_batch_id INT  NOT NULL,
                        wareHouse_id INT  NOT NULL,

                        FOREIGN KEY(employe_id) REFERENCES employees(id),
                        --FOREIGN KEY(product_batch_id) REFERENCES product_batches(id),
                        --FOREIGN KEY(wareHouse_id) REFERENCES warehouse(id),

                        UNIQUE(order_number)
);
