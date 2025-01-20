CREATE TABLE worker_resources
(
    id        BIGINT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    file_path VARCHAR(255),
    type      ENUM ('Identity Card', 'Police Certificate', 'Driver License', 'Payment'),
    worker_id BIGINT UNSIGNED,
    FOREIGN KEY (worker_id) REFERENCES workers (id)
)