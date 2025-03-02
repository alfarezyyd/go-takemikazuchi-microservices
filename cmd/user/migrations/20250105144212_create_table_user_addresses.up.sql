CREATE TABLE user_addresses
(
    id                     BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    place_id               VARCHAR(255),
    user_id                BIGINT UNSIGNED                            NOT NULL,
    formatted_address      TEXT,
    additional_information TEXT,
    street_number          VARCHAR(255),
    route                  VARCHAR(255),
    village                VARCHAR(255),
    district               VARCHAR(255),
    city                   VARCHAR(255),
    province               VARCHAR(255),
    country                VARCHAR(255),
    postal_code            VARCHAR(10),
    latitude               DECIMAL(9, 6),
    longitude              DECIMAL(9, 6),
    FOREIGN KEY (user_id) REFERENCES users (id)
)