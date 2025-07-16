DELETE FROM localities;
INSERT INTO localities (locality_name, province_name, country_name) VALUES
('Brooklyn', 'New York', 'United States'),
('Santa Monica', 'California', 'United States'),
('Cambridge', 'Massachusetts', 'United States'),
('Vancouver', 'British Columbia', 'Canada'),
('Medellín', 'Antioquia', 'Colombia');


DELETE FROM sellers;
INSERT INTO sellers (id, cid, company_name, address, telephone, locality_id) VALUES
(1, 1001, 'Alpha Traders Inc.', '123 Alpha St, New York, NY', '+1-212-555-0101', 1),
(2, 1008, 'Omicron Ventures', '888 Omicron Dr, San Francisco, CA', '+1-415-555-0110', 2),
(3, 1002, 'Beta Logistics Ltd.', '456 Beta Blvd, Chicago, IL', '+1-312-555-0102', 1),
(4, 1009, 'Lambda Freight Co.', '999 Lambda Way, Phoenix, AZ', '+1-602-555-0111', 2),
(5, 1003, 'Gamma Exports', '789 Gamma Ave, Houston, TX', '+1-713-555-0103', 3),
(6, 1004, 'Delta Wholesale', '321 Delta Rd, Miami, FL', '+1-305-555-0104', 3),
(7, 1005, 'Epsilon Products', '654 Epsilon Pkwy, Seattle, WA', '+1-206-555-0105', 4),
(8, 1006, 'Zeta Solutions', '987 Zeta Ln, Denver, CO', '+1-720-555-0106', 4),
(9, 1010, 'Sigma Technologies', '111 Sigma Blvd, Austin, TX', '+1-512-555-0112', 5),
(10, 1007, 'Theta Goods Co.', '246 Theta Cir, Boston, MA', '+1-617-555-0108', 5);
