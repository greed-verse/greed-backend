# Greed

## Overview

This project is a modular monolith designed to handle user accounts, wallets, and transactions using Go. The architecture combines synchronous communication interfaces and an in-memory event bus (powered by Watermill GoChannel) to provide a scalable and flexible system. The project aims to provide reliable wallet balance updates and transaction handling, ensuring security and performance.

## Features

- **User Management**:

  - Create, read, update, and delete users.
  - Enforce validations like unique email and username.

- **Wallet Management**:

  - Automatic wallet creation upon user registration.
  - Real-time balance updates on transactions.

- **Transactions**:

  - Secure fund transfers between wallets.
  - Atomic operations to ensure consistency.

- **Event-Driven Architecture**:

  - Asynchronous event handling using Watermill GoChannel.
  - Pub/sub model for decoupled module interaction.

## Technologies Used

- **Programming Language**: Go
- **Database**: PostgreSQL with `pgcrypto` for UUID generation.
- **Event Bus**: Watermill GoChannel for in-memory event-driven communication.
- **SQL Query Generation**: SQLC for type-safe query generation.
- **Dependency Injection**: Interfaces to decouple modules.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Project Structure](#project-structure)
3. [Database Schema](#database-schema)
4. [Event Workflow](#event-workflow)
5. [API Endpoints](#api-endpoints)
6. [Testing](#testing)
7. [Future Enhancements](#future-enhancements)

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL 13+
- Docker (optional, for containerized setup)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up the PostgreSQL database:
   ```bash
   psql -U postgres -c "CREATE DATABASE yourproject;"
   ```
4. Apply migrations (using SQLC):
   ```bash
   sqlc generate
   ```
5. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

```
project-root/
├── cmd/                # Application entry points
├── pkg/                # Core business logic
│   ├── account/        # Account module
│   ├── wallet/         # Wallet module
│   ├── transaction/    # Transaction module
│   ├── eventbus/       # Event bus (Watermill GoChannel)
│   └── shared/         # Shared utilities and interfaces
├── migrations/         # Database migrations
├── sqlc.yaml           # SQLC configuration
├── users.sql           # SQLC queries for users module
├── wallets.sql         # SQLC queries for wallets module
├── transactions.sql    # SQLC queries for transactions module
└── main.go             # Application entry point
```

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    CHECK (char_length(username) >= 3),
    CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);
```

### Wallets Table

```sql
CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    balance NUMERIC(12, 2) DEFAULT 0.00 NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
```

### Transactions Table

```sql
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_wallet_id UUID NOT NULL,
    receiver_wallet_id UUID NOT NULL,
    amount NUMERIC(12, 2) CHECK (amount > 0) NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    FOREIGN KEY (sender_wallet_id) REFERENCES wallets (id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_wallet_id) REFERENCES wallets (id) ON DELETE CASCADE
);
```

## Event Workflow

1. **User Registration**:

   - `account` module emits a `user.created` event.
   - `wallet` module subscribes and creates a wallet for the new user.

2. **Transaction Processing**:

   - A transaction is created in the `transactions` table.
   - Triggers ensure wallet balances are updated atomically.

## API Endpoints

### Users

- **POST /users**: Create a new user.
- **GET /users/{id}**: Get user details.

### Wallets

- **GET /wallets/{user_id}**: Get wallet details for a user.
- **POST /wallets/transfer**: Transfer funds between wallets.

### Transactions

- **GET /transactions/{id}**: Get transaction details.
- **POST /transactions**: Create a new transaction.

## Testing

- Run unit tests:
  ```bash
  go test ./...
  ```
- Test database integration:
  ```bash
  go test -tags=integration ./...
  ```

## Future Enhancements

1. Add support for multi-currency wallets.
2. Implement an external event bus (e.g., Kafka) for distributed systems.
3. Enhance security with JWT authentication.
4. Introduce rate limiting on transaction endpoints.

---

For any issues or suggestions, please open an issue or contribute via pull requests.

