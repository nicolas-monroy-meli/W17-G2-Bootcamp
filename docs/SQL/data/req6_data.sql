INSERT INTO buyers (id_card_number, first_name, last_name) VALUES
('CC1001', 'Juan', 'Pérez'),
('CC1002', 'María', 'González'),
('CC1003', 'Carlos', 'Ramírez'),
('CC1004', 'Ana', 'Martínez'),
('CC1005', 'Luis', 'Fernández'),
('CC1006', 'Sofía', 'Torres'),
('CC1007', 'Pedro', 'Ríos'),
('CC1008', 'Isabel', 'Cruz');

INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id) VALUES
('ORD-2024-001', '2024-06-10 08:30:00', 'TRK-501', 1),
('ORD-2024-002', '2024-06-11 09:15:00', 'TRK-502', 2),
('ORD-2024-003', '2024-06-11 10:50:00', 'TRK-503', 1),
('ORD-2024-004', '2024-06-12 14:20:00', 'TRK-504', 3),
('ORD-2024-005', '2024-06-13 13:00:00', 'TRK-505', 5),
('ORD-2024-006', '2024-06-14 15:30:00', 'TRK-506', 4),
('ORD-2024-007', '2024-06-15 11:48:00', 'TRK-507', 7),
('ORD-2024-008', '2024-06-16 16:43:00', 'TRK-508', 6),
('ORD-2024-009', '2024-06-17 09:50:00', 'TRK-509', 8);

INSERT INTO order_details (clean_liness_status, quantity, temperature, product_record_id, purchase_order_id) VALUES
('Limpio', 10, 12.5, NULL, 1),
('Sucia', 5, 15.0, NULL, 1),
('Limpio', 8, 13.0, NULL, 2),
('Limpio', 20, 11.5, NULL, 3),
('Limpio', 7, 15.2, NULL, 4),
('Sucia', 12, 16.0, NULL, 4),
('Limpio', 18, 10.7, NULL, 5),
('Limpio', 24, 14.8, NULL, 6),
('Limpio', 11, 11.2, NULL, 7),
('Sucia', 6, 15.3, NULL, 7),
('Limpio', 15, 12.1, NULL, 8),
('Limpio', 9, 13.9, NULL, 8),
('Limpio', 13, 14.5, NULL, 9),
('Sucia', 4, 17.0, NULL, 9),
('Limpio', 17, 10.9, NULL, 5);