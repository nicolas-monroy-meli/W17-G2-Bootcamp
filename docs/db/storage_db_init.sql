USE `storage_db`;

CREATE TABLE `sections` (
                         `id` INT AUTO_INCREMENT PRIMARY KEY,
                         `sectionNumber` INT NOT NULL CHECK (sectionNumber > 0),
                         `currentTemperature` DOUBLE NOT NULL,
                         `minimumTemperature` DOUBLE NOT NULL,
                         `currentCapacity` INT NOT NULL,
                         `minimumCapacity` INT NOT NULL CHECK (minimumCapacity > 0),
                         `maximumCapacity` INT NOT NULL,
                         `warehouseID` INT NOT NULL CHECK (warehouseID >= 1),
                         `productTypeID` INT NOT NULL CHECK (productTypeID >= 1),

    -- Cross-field constraints
                         CHECK (currentTemperature >= minimumTemperature),
                         CHECK (currentCapacity >= minimumCapacity AND currentCapacity <= maximumCapacity),
                         CHECK (maximumCapacity > minimumCapacity)

    -- Add FOREIGN KEY lines if you have the referenced tables:
    -- ,FOREIGN KEY (warehouseID) REFERENCES Warehouse(id)
    -- ,FOREIGN KEY (productTypeID) REFERENCES ProductType(id)
);

-- Optional: add indexes if you'll filter/search by fields like warehouseID/productTypeID
-- CREATE INDEX idx_section_warehouse ON Section (warehouseID);
-- CREATE INDEX idx_section_product_type ON Section (productTypeID);