CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- Unique wallet ID (UUID)
    user_id UUID NOT NULL,                               -- Foreign key linking to users
    balance NUMERIC(12, 2) DEFAULT 0.00 NOT NULL,        -- Wallet balance with 2 decimal precision
    updated_at timestamptz DEFAULT NOW() NOT NULL, -- Last wallet update timestamp
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE -- Link to users table
);

-- Indexes for performance optimization
CREATE INDEX idx_wallets_user_id ON wallets (user_id); -- Index for wallet-to-user relationship
CREATE INDEX idx_wallets_balance ON wallets (balance); -- Index for balance-related queries
