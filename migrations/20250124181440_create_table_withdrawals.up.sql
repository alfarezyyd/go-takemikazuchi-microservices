CREATE TABLE withdrawals
(
    id               BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    worker_id        BIGINT UNSIGNED                            NOT NULL,                           -- ID worker yang mencairkan dana
    wallet_id        BIGINT UNSIGNED                            NOT NULL,                           -- ID rekening/e-wallet tujuan pencairan
    amount           DECIMAL(15, 2)                             NOT NULL,                           -- Jumlah dana yang dicairkan
    status           ENUM ('Pending', 'Approved', 'Rejected')   NOT NULL DEFAULT 'Pending',         -- Status pencairan
    requested_at     TIMESTAMP                                           DEFAULT CURRENT_TIMESTAMP, -- Waktu request pencairan
    processed_at     TIMESTAMP                                  NULL,                               -- Waktu pencairan diproses (null jika masih pending)
    admin_id         BIGINT UNSIGNED                            NULL,                               -- Admin yang memproses pencairan (null jika belum diproses)
    rejection_reason TEXT                                       NULL,                               -- Alasan penolakan jika statusnya 'Rejected'

    -- Relasi
    FOREIGN KEY (worker_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (wallet_id) REFERENCES worker_wallets (id) ON DELETE CASCADE,
    FOREIGN KEY (admin_id) REFERENCES users (id) ON DELETE SET NULL
);
