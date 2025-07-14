DELETE FROM localities;

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