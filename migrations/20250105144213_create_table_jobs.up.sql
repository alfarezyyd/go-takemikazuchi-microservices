CREATE TABLE jobs
(
    id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id     BIGINT UNSIGNED         NOT NULL,
    category_id BIGINT UNSIGNED         NOT NULL,
    title       VARCHAR(255)            NOT NULL,
    description TEXT                    NOT NULL,
    latitude    DECIMAL(9, 6),
    longitude   DECIMAL(9, 6),
    address     TEXT,
    place_id    VARCHAR(255),
    price       DECIMAL(10, 2) UNSIGNED NOT NULL,
    status      ENUM ('Open', 'Closed', 'Done') DEFAULT 'Open',
    created_at  TIMESTAMP                       DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
