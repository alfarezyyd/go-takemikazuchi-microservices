CREATE TABLE worker_resources
(
    id         BIGINT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    image_path VARCHAR(255),
    video_url  VARCHAR(255),
    worker_id  BIGINT UNSIGNED,
    FOREIGN KEY (worker_id) REFERENCES workers (id)
)