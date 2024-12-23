CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    profile_image TEXT,
    created_at timestamptz DEFAULT NOW()
);

-- Create an index on the email field for faster lookups
CREATE UNIQUE INDEX idx_users_email ON users(email);

-- Create an index on the id field for faster lookups
CREATE UNIQUE INDEX idx_users_id ON users(id);
