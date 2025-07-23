DROP TABLE IF EXISTS sections;
DROP TABLE IF EXISTS product_batches;

CREATE TABLE `sections` (
                            `id` INT AUTO_INCREMENT PRIMARY KEY,
                            `section_number` INT unique NOT NULL CHECK (section_number > 0),
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
    -- ,FOREIGN KEY (warehouse_id) REFERENCES Warehouse(id)
    -- ,FOREIGN KEY (product_type_id) REFERENCES ProductType(id)
);
-- CREATE INDEX idx_section_warehouse ON Section (warehouseID);
-- CREATE INDEX idx_section_product_type ON Section (productTypeID);

CREATE TABLE `product_batches` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `batch_number` int unique NOT NULL,
                                   `current_quantity` int NOT NULL,
                                   `initial_quantity` int NOT NULL,
                                   `current_temperature` int NOT NULL,
                                   `minimum_temperature` int NOT NULL,
                                   `due_date` datetime NOT NULL,
                                   `manufacturing_date` date NOT NULL,
                                   `manufacturing_hour` varchar(12) NOT NULL,
                                   `product_id` int NOT NULL,
                                   `section_id` int NOT NULL,
                                   PRIMARY KEY (`id`),
                                   KEY `fk_section_id` (`section_id`),
                                   CONSTRAINT `fk_section_id` FOREIGN KEY (`section_id`) REFERENCES `sections` (`id`),
                                   CONSTRAINT `product_batches_chk_1` CHECK ((`batch_number` > 0)),
                                   CONSTRAINT `product_batches_chk_2` CHECK ((`initial_quantity` > 0)),
                                   CONSTRAINT `product_batches_chk_3` CHECK ((`product_id` > 0)),
                                   CONSTRAINT `product_batches_chk_4` CHECK ((`section_id` > 0)),
                                   CONSTRAINT `product_batches_chk_5` CHECK ((`current_quantity` >= `initial_quantity`)),
                                   CONSTRAINT `product_batches_chk_6` CHECK ((`current_temperature` >= `minimum_temperature`))
);
