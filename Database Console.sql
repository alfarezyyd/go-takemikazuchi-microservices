DROP DATABASE IF EXISTS go_takemikazuchi_microservices;
CREATE DATABASE go_takemikazuchi_microservices;
USE go_takemikazuchi_microservices;

SELECT * FROM workers;
SELECT * FROM jobs;
SELECT * FROM job_resources;
SELECT * FROM job_applications;
SELECT * FROM worker_resources;
SELECT * FROM worker_wallets;
SELECT * FROM user_addresses;
SELECT * FROM users;
SELECT * FROM withdrawals;
SELECT * FROM orders;
SELECT * FROM transactions;
UPDATE users SET role = 'Admin' WHERE id = 1;
DELETE FROM workers;
DELETE FROM worker_resources;
DELETE FROM worker_wallets;
DELETE FROM transactions;


DROP TABLE IF EXISTS reviews;

