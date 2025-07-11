USE frescos_db;
SET SQL_SAFE_UPDATES = 0;

DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_code VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    height DOUBLE NOT NULL CHECK (height >= 0),
    length DOUBLE NOT NULL CHECK (length >= 0),
    width DOUBLE NOT NULL CHECK (width >= 0),
    net_weight DOUBLE NOT NULL CHECK (net_weight >= 0),
    expiration_rate DOUBLE NOT NULL CHECK (expiration_rate >= 0),
    freezing_rate DOUBLE NOT NULL,
    recommended_freezing_temperature DOUBLE NOT NULL,
    product_type_id INT UNSIGNED NOT NULL,
    seller_id INT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    -- Foraneas
    -- FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    -- FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

DROP TABLE IF EXISTS product_records;
CREATE TABLE product_records (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    last_update_date DATETIME(6) NOT NULL,
    purchase_price DECIMAL(19,2) NOT NULL,
    sale_price DECIMAL(19,2) NOT NULL,
    product_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);