CREATE TABLE job_resources
(
    id         BIGINT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
    image_path VARCHAR(255),
    video_url  VARCHAR(255),
    job_id     BIGINT UNSIGNED,
    FOREIGN KEY (job_id) REFERENCES jobs (id)
)