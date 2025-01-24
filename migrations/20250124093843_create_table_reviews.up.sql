CREATE TABLE reviews
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    reviewer_id BIGINT UNSIGNED                            NOT NULL, -- ID user yang memberi review (bisa employer/worker)
    reviewed_id BIGINT UNSIGNED                            NOT NULL, -- ID user yang menerima review
    job_id      BIGINT UNSIGNED                            NOT NULL,
    role        ENUM ('Worker', 'Employer')                NOT NULL, -- Menentukan siapa yang direview
    rating      TINYINT UNSIGNED                           NOT NULL CHECK (rating BETWEEN 1 AND 5),
    review_text TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- Relasi
    FOREIGN KEY (reviewer_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (reviewed_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (job_id) REFERENCES jobs (id) ON DELETE CASCADE
);
