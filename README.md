# Transaction Processing Engine

## Objective
A backend service that processes card transactions and maintains card balances, simulating a simplified payment switch authorization engine.

## Architecture Layers

The application is structured into clearly defined layers to separate concerns, making the codebase scalable, testable, and maintainable.

### 🔹 1. Handler Layer (`handler/`)
- Handles HTTP requests & responses.
- Parses input (JSON, URL params).
- Calls service layer.
- Returns formatted JSON response.

👉 Keeps HTTP logic separate from business logic.

---

### 🔹 2. Service Layer (`service/`)
- Core business logic.
- Validates:
  - Card existence & status
  - PIN verification (SHA256)
  - Transaction rules
- Processes:
  - Withdraw / Top-up
- Logs every transaction (success + failure).

👉 This is the **brain of the system**.

---

### 🔹 3. Repository Layer (`repository/`)
- Acts as an abstraction over storage.
- Provides clean data access methods.
- Decouples service from storage implementation.

👉 Makes system scalable (easy to switch DB).

---

### 🔹 4. Storage Layer (`storage/`)
- In-memory database using Go maps.
- Thread-safe using mutex.
- Stores:
  - Cards
  - Transactions

👉 Simulates database behavior.

---

### 🔹 5. Model (`model/`) & Utilities (`utils/`)
- **Model**: Defines data domains like `Card`, `Transaction`, API schemas, and response codes.
- **Utils**: Contains helper logic such as SHA-256 password hashing.

## Setup Instructions

### Prerequisites
- [Go (Golang)](https://golang.org/dl/) 1.18+
- [Git](https://git-scm.com/)

### Clone the Repository
```bash
git clone https://github.com/RuneRogue/Transaction-Processing-Engine.git
cd Transaction-Processing-Engine
```

### Install Dependencies
```bash
# This fetches required packages like chi for routing, cors, and godotenv.
go mod tidy
```

## Run Steps

### Starting the Server
Run the application directly:
```bash
go run main.go
```
By default, the server starts on port `8080`. You can configure a custom port via the `.env` file (e.g., `PORT=8080`).

## API Examples (cURL)

### 1. Perform a Transaction

**Withdraw Example:**
```bash
curl -X POST http://localhost:8080/api/transaction \
     -H "Content-Type: application/json" \
     -d '{
           "cardNumber": "4123456789012345",
           "pin": "1234",
           "type": "withdraw",
           "amount": 200
         }'
```

**Top-up Example:**
```bash
curl -X POST http://localhost:8080/api/transaction \
     -H "Content-Type: application/json" \
     -d '{
           "cardNumber": "4123456789012345",
           "pin": "1234",
           "type": "topup",
           "amount": 500
         }'
```

**Expected Success Response:**
```json
{
  "status": "SUCCESS",
  "respCode": "00",
  "balance": 800
}
```

**Expected Insufficient Balance Response:**
```json
{
  "status": "FAILED",
  "respCode": "99",
  "message": "Insufficient balance"
}
```

### 2. Get Card Balance

**Request:**
```bash
curl -X GET http://localhost:8080/api/card/balance/4123456789012345
```

### 3. Get Transaction History

**Request:**
```bash
curl -X GET http://localhost:8080/api/card/transactions/4123456789012345
```
