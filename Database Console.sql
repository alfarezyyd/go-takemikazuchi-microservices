DROP DATABASE IF EXISTS go_takemikazuchi_microservices;
CREATE DATABASE go_takemikazuchi_microservices;
USE go_takemikazuchi_microservices;

SELECT *
FROM workers;
SELECT *
FROM jobs;
SELECT *
FROM job_resources;
SELECT *
FROM job_applications;
SELECT *
FROM worker_resources;
SELECT *
FROM worker_wallets;
UPDATE worker_wallets SET worker_id = 28 WHERE worker_id = 10;
SELECT *
FROM user_addresses;
SELECT *
FROM users;
SELECT *
FROM withdrawals;
SELECT *
FROM orders;
SELECT *
FROM transactions;
UPDATE users
SET role = 'Admin'
WHERE id = 1;
DELETE
FROM workers;
DELETE
FROM worker_resources;
DELETE
FROM worker_wallets;
DELETE
FROM transactions;


DROP TABLE IF EXISTS reviews;

CREATE DATABASE go_takemikazuchi_microservices_users;
CREATE DATABASE go_takemikazuchi_microservices_jobs;
CREATE DATABASE go_takemikazuchi_microservices_workers;
CREATE DATABASE go_takemikazuchi_microservices_payments;
DROP DATABASE go_takemikazuchi_microservices_workers;
DROP DATABASE go_takemikazuchi_microservices_users;
DROP DATABASE go_takemikazuchi_microservices_payments;
USE go_takemikazuchi_microservices_users;
USE go_takemikazuchi_microservices_categories;
USE go_takemikazuchi_microservices_jobs;
USE go_takemikazuchi_microservices_workers;
USE go_takemikazuchi_microservices_payments;

SELECT * FROM jobs;
SELECT * FROM user_addresses;
SELECT * FROM categories;
SELECT * FROM transactions;

SELECT *
FROM jobs;

SELECT *
FROM job_applications;

SELECT *
FROM categories;
SELECT *
FROM users;
UPDATE users
SET role = 'Admin'
WHERE id = 27;
SELECT *
FROM users;

DELETE
FROM one_time_password_tokens;
DELETE
FROM users;

DELETE FROM user_addresses;
DELETE FROM one_time_password_tokens;
DELETE FROM users;
