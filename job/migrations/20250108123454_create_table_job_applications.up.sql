CREATE TABLE job_applications
(
    id           BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    job_id       BIGINT UNSIGNED                            NOT NULL,
    applicant_id BIGINT UNSIGNED                            NOT NULL,
    status       ENUM ('Pending', 'Rejected', 'Accepted') DEFAULT 'Pending',
    created_at   TIMESTAMP                                DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP                                DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs (id)
);