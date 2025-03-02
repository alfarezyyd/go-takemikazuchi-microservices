CREATE TABLE users
(
    id                BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name              VARCHAR(100)                               NOT NULL,
    email             VARCHAR(255) UNIQUE                        NOT NULL,
    password          VARCHAR(255)                               NOT NULL,
    role              ENUM ('Worker', 'Employer', 'Admin')       NOT NULL DEFAULT 'Employer',
    phone_number      VARCHAR(15),
    profile_picture   TEXT,
    is_active         BOOLEAN                                             DEFAULT TRUE,
    created_at        TIMESTAMP                                           DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP                                           DEFAULT CURRENT_TIMESTAMP,
    email_verified_at TIMESTAMP
);
