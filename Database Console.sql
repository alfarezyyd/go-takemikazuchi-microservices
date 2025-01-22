DROP DATABASE IF EXISTS go_takemikazuchi_api;
CREATE DATABASE go_takemikazuchi_api;
USE go_takemikazuchi_api;

SELECT * FROM workers;
SELECT * FROM jobs;
SELECT * FROM job_applications;
SELECT * FROM worker_resources;
SELECT * FROM worker_wallets;
SELECT * FROM user_addresses;
SELECT * FROM users;
UPDATE users SET role = 'Admin' WHERE id = 1;
DELETE FROM workers;
DELETE FROM worker_resources;
DELETE FROM worker_wallets;