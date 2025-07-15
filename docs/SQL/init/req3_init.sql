USE `frescos_db`;

DROP TABLE IF EXISTS sections;
CREATE TABLE `sections` (
                            `id` INT AUTO_INCREMENT PRIMARY KEY,
                            `section_number` INT NOT NULL CHECK (section_number > 0),
                            `current_temperature` DOUBLE NOT NULL,
                            `minimum_temperature` DOUBLE NOT NULL,
                            `current_capacity` INT NOT NULL,
                            `minimum_capacity` INT NOT NULL CHECK (minimum_capacity > 0),
                            `maximum_capacity` INT NOT NULL,
                            `warehouse_id` INT NOT NULL CHECK (warehouse_id >= 1),
                            `product_type_id` INT NOT NULL CHECK (product_type_id >= 1),

    -- Cross-field constraints
                            CHECK (current_temperature >= minimum_temperature),
                            CHECK (current_capacity >= minimum_capacity AND current_capacity <= maximum_capacity),
                            CHECK (maximum_capacity > minimum_capacity)

    -- Add FOREIGN KEY lines if you have the referenced tables:
    -- ,FOREIGN KEY (warehouse_id) REFERENCES Warehouse(id)
    -- ,FOREIGN KEY (product_type_id) REFERENCES ProductType(id)
);

-- Optional: add indexes if you'll filter/search by fields like warehouseID/productTypeID
-- CREATE INDEX idx_section_warehouse ON Section (warehouseID);
-- CREATE INDEX idx_section_product_type ON Section (productTypeID);

CREATE TABLE ProductBatch (
                              id                  INT AUTO_INCREMENT PRIMARY KEY,
                              batch_number         INT NOT NULL CHECK (batch_number > 0),
                              current_quantity     INT NOT NULL,
                              initial_quantity     INT NOT NULL CHECK (initial_quantity > 0),
                              current_temperature  INT NOT NULL,
                              minimum_temperature  INT NOT NULL,
                              due_date             DATETIME NOT NULL,
                              manufacturing_date   DATE NOT NULL,
                              manufacturing_hour   TIME NOT NULL,
                              product_id           INT NOT NULL CHECK (product_id > 0),
                              section_id           INT NOT NULL CHECK (section_id > 0),

    -- Cross-field constraints
                              CHECK (current_quantity >= initial_quantity),
                              CHECK (current_temperature >= minimum_temperature)

    -- Foreign keys (descomenta si tienes las tablas referenciadas)
    -- ,FOREIGN KEY (product_id) REFERENCES Product(id)
    -- ,FOREIGN KEY (section_id) REFERENCES Section(id)
);
