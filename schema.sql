-- Tabel Users (Akun Pengguna & Admin)
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'user', -- 'user' atau 'admin'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Accounts (Penyimpanan Saldo / Dompet Digital User)
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (user_id) ON DELETE CASCADE,
    balance DECIMAL(15, 2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Campaign (Galang Dana / Program Donasi)
CREATE TABLE campaigns (
    campaign_id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    target_amount DECIMAL(15, 2) NOT NULL,
    current_amount DECIMAL(15, 2) DEFAULT 0.00,
    status VARCHAR(20) DEFAULT 'active', -- 'active', 'completed', 'closed'
    created_by INT REFERENCES users (user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Payment Methods (Metode Pembayaran yang Tersedia)
CREATE TABLE payment_methods (
    payment_method_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL, -- Contoh: BCA, QRIS, GoPay, OVO
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Donations (Donasi yang Masuk ke Campaign)
CREATE TABLE donations (
    donation_id SERIAL PRIMARY KEY,
    campaign_id INT REFERENCES campaigns (campaign_id) ON DELETE CASCADE,
    user_id INT REFERENCES users (user_id) ON DELETE SET NULL,
    amount DECIMAL(15, 2) NOT NULL,
    comment TEXT,
    anonymous BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Transactions (Pencatatan Alur Transaksi Keuangan)
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    donation_id INT REFERENCES donations (donation_id) ON DELETE CASCADE,
    payment_method_id INT REFERENCES payment_methods (payment_method_id),
    status VARCHAR(30) DEFAULT 'pending', -- 'pending', 'success', 'failed'
    external_id VARCHAR(100), -- ID dari Payment Gateway (misal: Midtrans/Xendit)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexing untuk optimasi query
CREATE INDEX idx_campaigns_status ON campaigns (status);

CREATE INDEX idx_donations_campaign ON donations (campaign_id);