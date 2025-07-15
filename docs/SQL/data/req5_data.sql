INSERT INTO employees (id_card_number, first_name, last_name, wareHouse_id)
VALUES ('1001', 'Juan', 'Perez', 1),
       ('1002', 'Maria', 'Gomez', 2),
       ('1003', 'Carlos', 'Rodriguez', 1),
       ('1004', 'Ana', 'Lopez', 3),
       ('1005', 'Pedro', 'Martinez', 2),
       ('1006', 'Laura', 'Sanchez', 1),
       ('1007', 'David', 'Garcia', 3),
       ('1008', 'Sofia', 'Ramirez', 2),
       ('1009', 'Pablo', 'Torres', 1),
       ('1010', 'Elena', 'Diaz', 3);

INSERT INTO inbound_orders (order_date, order_number, employe_id, product_batch_id, wareHouse_id)
VALUES ('2023-01-15', 'ORD-2023-001', 1, 101, 1),
       ('2023-01-16', 'ORD-2023-002', 2, 102, 2),
       ('2023-01-17', 'ORD-2023-003', 3, 103, 1),
       ('2023-01-18', 'ORD-2023-004', 4, 104, 3),
       ('2023-01-19', 'ORD-2023-005', 5, 105, 2),
       ('2023-01-20', 'ORD-2023-006', 6, 106, 1),
       ('2023-01-21', 'ORD-2023-007', 7, 107, 3),
       ('2023-01-22', 'ORD-2023-008', 8, 108, 2),
       ('2023-01-23', 'ORD-2023-009', 9, 109, 1),
       ('2023-01-24', 'ORD-2023-010', 10, 110, 3);