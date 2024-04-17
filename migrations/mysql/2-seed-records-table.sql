-- +migrate Up
INSERT INTO records (name, marks, created_at) VALUES
('John Doe', '[100,200,300]', NOW()),
('Jane Doe', '[50,150,250]', NOW()),
('Remon Di', '[50,50,50]', NOW()),
('Mark Fan', '[30,45,75]', NOW()),
('Bob Sila', '[100,0,10]', NOW()),
('Pahn Sei', '[7,50,50]', NOW());