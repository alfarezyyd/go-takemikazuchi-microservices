CREATE TABLE worker_skills
(
    id               BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    worker_id        BIGINT UNSIGNED                            NOT NULL,
    skill_id         BIGINT UNSIGNED                            NOT NULL,
    experience_level ENUM ('Beginner', 'Intermediate', 'Expert') DEFAULT 'Beginner', -- Opsional
    FOREIGN KEY (worker_id) REFERENCES workers (id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE
);