CREATE TABLE reviews
(
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    reviewer_id BIGINT UNSIGNED NOT NULL,
    reviewee_id BIGINT UNSIGNED NOT NULL,
    job_id      BIGINT UNSIGNED NOT NULL,
    rating      DECIMAL(2, 1)   NOT NULL,
    comment     TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (reviewer_id) REFERENCES users (id),
    FOREIGN KEY (reviewee_id) REFERENCES users (id),
    FOREIGN KEY (job_id) REFERENCES jobs (id)
);
