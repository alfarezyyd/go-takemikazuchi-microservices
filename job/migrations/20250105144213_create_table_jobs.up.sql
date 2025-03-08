CREATE TABLE jobs
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id     BIGINT UNSIGNED         NOT NULL,
    address_id  BIGINT UNSIGNED         NOT NULL,
    category_id BIGINT UNSIGNED         NOT NULL,
    worker_id   BIGINT UNSIGNED         NOT NULL,
    title       VARCHAR(255)            NOT NULL,
    description TEXT                    NOT NULL,
    price       DECIMAL(10, 2) UNSIGNED NOT NULL,
    status      ENUM ('Open', 'Process', 'On Working','Closed', 'Done') DEFAULT 'Open',
    created_at  TIMESTAMP                                               DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                                               DEFAULT CURRENT_TIMESTAMP
);
