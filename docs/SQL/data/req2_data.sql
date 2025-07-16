DELETE FROM localities;

DELETE FROM warehouses;


INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES
('WH-001', '789 Industrial Ave, Brooklyn, NY', '+1-718-555-0201', 2500, -18.5),
('WH-002', '456 Tech Valley Rd, Santa Monica, CA', '+1-310-555-0202', 1800, 15.0),
('WH-003', '321 Innovation Parkway, Cambridge, MA', '+1-617-555-0203', 3000, 2.0),
('WH-004', '654 Portside Drive, Vancouver', '+1-604-555-0204', 4200, -25.0),
('WH-005', '987 Coffee District, Medellín', '+1-574-555-0205', 1500, 18.0),
('WH-006', '135 Automotive Blvd, Brooklyn, NY', '+1-718-555-0206', 3500, 10.0),
('WH-007', '246 Pharma Park, Santa Monica, CA', '+1-310-555-0207', 2000, 5.0),
('WH-008', '864 Agro Center, Cambridge, MA', '+1-617-555-0208', 2800, 12.5),
('WH-009', '753 Cold Storage Lane, Vancouver', '+1-604-555-0209', 5000, -30.0),
('WH-010', '369 Textile Complex, Medellín', '+1-574-555-0210', 3200, 20.0);

INSERT INTO localities (id, name) VALUES
(1, 'Brooklyn'),
(2, 'Santa Monica'),
(3, 'Cambridge'),
(4, 'Vancouver'),
(5, 'Medellín');

DELETE FROM carries;

INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES
('1001', 'Alpha Traders Inc.', '123 Alpha St, New York, NY', '+1-212-555-0101', 1),
('1008', 'Omicron Ventures', '888 Omicron Dr, San Francisco, CA', '+1-415-555-0110', 2),
('1002', 'Beta Logistics Ltd.', '456 Beta Blvd, Chicago, IL', '+1-312-555-0102', 1),
('1009', 'Lambda Freight Co.', '999 Lambda Way, Phoenix, AZ', '+1-602-555-0111', 2),
('1003', 'Gamma Exports', '789 Gamma Ave, Houston, TX', '+1-713-555-0103', 3),
('1004', 'Delta Wholesale', '321 Delta Rd, Miami, FL', '+1-305-555-0104', 3),
('1005', 'Epsilon Products', '654 Epsilon Pkwy, Seattle, WA', '+1-206-555-0105', 4),
('1006', 'Zeta Solutions', '987 Zeta Ln, Denver, CO', '+1-720-555-0106', 4),
('1010', 'Sigma Technologies', '111 Sigma Blvd, Austin, TX', '+1-512-555-0112', 5),
('1007', 'Theta Goods Co.', '246 Theta Cir, Boston, MA', '+1-617-555-0108', 5);