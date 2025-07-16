DELETE FROM sections;
DELETE FROM product_batches;

INSERT INTO `sections` (
    `section_number`,
    `current_temperature`,
    `minimum_temperature`,
    `current_capacity`,
    `minimum_capacity`,
    `maximum_capacity`,
    `warehouse_id`,
    `product_type_id`) VALUES
                           (1, 0, -5, 50, 20, 100, 1, 1),
                           (2, -2, -6, 60, 30, 110, 2, 2),
                           (3, 1, -4, 70, 40, 120, 3, 3),
                           (4, -3, -7, 80, 50, 130, 4, 4),
                           (5, 2, -5, 90, 60, 140, 5, 5),
                           (6, -4, -8, 100, 70, 150, 6, 6),
                           (7, 3, -6, 110, 80, 160, 7, 7),
                           (8, -5, -9, 120, 90, 170, 8, 8),
                           (9, 4, -7, 130, 100, 180, 9, 9),
                           (10, -6, -10, 140, 110, 190, 10, 10);

INSERT INTO `product_batches` (
    batch_number,
    current_quantity,
    initial_quantity,
    current_temperature,
    minimum_temperature,
    due_date,
    manufacturing_date,
    manufacturing_hour,
    product_id,
    section_id
) VALUES
      (1, 200, 200, 2, -5, '2024-07-05 17:00:00', '2024-06-01', '08:00:00', 1, 1),
      (2, 310, 310, -2, -6, '2024-08-01 12:00:00', '2024-07-01', '09:30:00', 2, 2),
      (3, 400, 400, 3, -4, '2024-07-15 08:30:00', '2024-05-10', '10:30:00', 3, 3),
      (4, 500, 500, -3, -7, '2024-09-10 16:30:00', '2024-06-20', '14:45:00', 4, 4),
      (5, 600, 600, 2, -5, '2024-10-09 10:00:00', '2024-08-18', '07:15:00', 5, 5);4, 80, 80, 3, -6, '2024-08-10 11:00:00', '2024-06-10', '10:00:00', 4, 7);