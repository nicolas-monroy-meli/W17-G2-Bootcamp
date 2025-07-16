DELETE FROM products;
INSERT INTO products (
    product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id
) VALUES
('TOM123', 'Caja de tomate fresco', 20.5, 30.2, 15.0, 8.2, 1.0, 0.0, -1.0, 1, 1),
('LECH456', 'Lechuga orgánica', 8.3, 12.0, 8.0, 0.5, 0.8, 0.0, -0.5, 2, 2),
('MANZ789', 'Manzana roja premium', 7.0, 10.0, 8.0, 0.45, 1.2, 0.0, -0.5, 3, 1),
('ZAN901', 'Zanahoria seleccionada', 18.0, 25.0, 5.5, 1.1, 0.9, 0.0, -0.7, 4, 3),
('PAPA321', 'Papa blanca lavada', 12.0, 18.0, 12.0, 3.8, 1.1, 0.0, -1.0, 1, 2);


DELETE FROM product_records;
INSERT INTO product_records (
    last_update_date, purchase_price, sale_price, product_id
) VALUES
('2024-06-15 10:23:00.000000', 10.00, 15.50, 1),
('2024-06-16 08:00:00.000000', 4.00, 6.50, 2);