CREATE TABLE currencies (id serial primary key , code varchar(50), full_name varchar(50), sign varchar(10), to_dollar decimal(6,4));

INSERT INTO currencies (code, full_name, sign, to_dollar) VALUES
('USD', 'United States Dollar', '$', 1.0000),
('EUR', 'Euro', '€', 1.1000),
('JPY', 'Japanese Yen', '¥', 0.0075),
('GBP', 'British Pound', '£', 1.2500),
('AUD', 'Australian Dollar', '$', 0.6800),
('CAD', 'Canadian Dollar', '$', 0.7400),
('CHF', 'Swiss Franc', 'CHF', 1.0900),
('CNY', 'Chinese Yuan', '¥', 0.1400),
('SEK', 'Swedish Krona', 'kr', 0.0950),
('NZD', 'New Zealand Dollar', '$', 0.6200),
('MXN', 'Mexican Peso', '$', 0.0550),
('SGD', 'Singapore Dollar', '$', 0.7400),
('HKD', 'Hong Kong Dollar', '$', 0.1270),
('NOK', 'Norwegian Krone', 'kr', 0.0900),
('KRW', 'South Korean Won', '₩', 0.00075);
