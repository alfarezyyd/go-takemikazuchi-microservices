CREATE TABLE skills
(
    id   BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(100) UNIQUE                        NOT NULL -- Nama skill, misalnya "Bersih-bersih", "Angkut Barang"
);