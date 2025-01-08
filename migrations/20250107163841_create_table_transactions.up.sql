CREATE TABLE transactions
(
    id             VARCHAR(255) PRIMARY KEY,                                        -- ID transaksi unik
    job_id         BIGINT UNSIGNED         NOT NULL,                                -- ID pekerjaan terkait
    payer_id       BIGINT UNSIGNED         NOT NULL,                                -- ID pemberi kerja (pembayar)
    payee_id       BIGINT UNSIGNED         NOT NULL,                                -- ID pekerja (penerima)
    amount         DECIMAL(10, 2) UNSIGNED NOT NULL,                                -- Jumlah transaksi
    snap_token     VARCHAR(255),
    payment_method VARCHAR(50)             NOT NULL,                                -- Metode pembayaran (e.g., 'credit_card', 'bank_transfer', 'cash')
    status         ENUM ('Pending', 'Completed', 'Failed')                          -- Status transaksi
                             DEFAULT 'Pending',
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                             -- Tanggal transaksi
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update terakhir
    FOREIGN KEY (job_id) REFERENCES jobs (id),
    FOREIGN KEY (payer_id) REFERENCES users (id),
    FOREIGN KEY (payee_id) REFERENCES users (id)
);
