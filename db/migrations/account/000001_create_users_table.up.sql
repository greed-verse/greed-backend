CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),     -- Unique identifier (UUID)
    username VARCHAR(50) NOT NULL UNIQUE,             -- Unique username
    email VARCHAR(255) NOT NULL UNIQUE,               -- Unique email address
    created_at timestamptz DEFAULT NOW() NOT NULL, -- Account creation timestamp
    CHECK (char_length(username) >= 3)               -- Ensure username has at least 3 characters
);

-- Indexes for fast lookups
CREATE INDEX idx_users_id ON users (id);            -- Index on the primary key
CREATE INDEX idx_users_email ON users (email);      -- Index for efficient email lookups
