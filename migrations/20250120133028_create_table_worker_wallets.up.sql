CREATE TABLE worker_wallets
(
    id             BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT                    NOT NULL,
    worker_id      BIGINT UNSIGNED                                               NOT NULL,
    wallet_type    ENUM ('Bank', 'DANA', 'OVO', 'GoPay', 'LinkAja', 'ShopeePay') NOT NULL,
    account_name   VARCHAR(100)                                                  NOT NULL, -- Nama pemilik rekening/e-wallet
    account_number VARCHAR(50)                                                   NOT NULL, -- Nomor rekening/e-wallet
    bank_name      VARCHAR(100),                                                           -- Hanya diisi jika wallet_type = 'Bank'
    is_primary     BOOLEAN   DEFAULT FALSE,                                                -- Menandai akun utama untuk pembayaran
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (worker_id) REFERENCES workers (id)
);