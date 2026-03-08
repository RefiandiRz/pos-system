# 🛒 POS System — Point of Sale Backend

A RESTful Point of Sale (POS) backend API built with **Go**, **Fiber v3**, **GORM**, and **PostgreSQL**. Designed as a portfolio project showcasing clean architecture, JWT authentication, and Docker containerization.

---

## 🚀 Tech Stack

| Layer        | Technology          |
|--------------|---------------------|
| Language     | Go 1.25+            |
| Framework    | Fiber v3            |
| Database     | PostgreSQL 16       |
| ORM          | GORM                |
| Auth         | JWT                 |
| Container    | Docker & Docker Compose |

---

## 📦 Features

- 🔐 JWT Authentication with role-based access (Admin / Cashier)
- 📦 Product & Category management
- 🧾 Sales / Checkout flow with auto stock deduction
- 📊 Sales reports & top product analytics

---

## 🗂️ Project Structure

```
pos-system/
├── cmd/                    # Entry point
├── config/                 # Database connection
├── internal/
│   ├── handlers/           # HTTP handlers
│   ├── services/           # Business logic
│   ├── repositories/       # Database queries
│   ├── models/             # GORM models
│   └── middleware/         # JWT middleware
├── routes/                 # Route definitions
├── utils/                  # JWT helpers
├── Dockerfile
├── docker-compose.yml
└── .env.example
```

---

## ⚙️ Getting Started

### Prerequisites
- [Go 1.25+](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/pos-system.git
cd pos-system
```

### 2. Setup environment variables
```bash
cp .env.example .env
# Edit .env with your own values
```

### 3. Start with Docker
```bash
docker compose up -d
```

### 4. Run the app locally (without Docker app container)
```bash
docker compose up -d postgres
go run cmd/main.go
```

---

## 🌐 API Endpoints

### Auth
| Method | Endpoint | Access |
|--------|----------|--------|
| POST | `/api/auth/register` | Public |
| POST | `/api/auth/login` | Public |

### Products
| Method | Endpoint | Access |
|--------|----------|--------|
| GET | `/api/products` | All |
| POST | `/api/products` | Admin |
| PUT | `/api/products/:id` | Admin |
| DELETE | `/api/products/:id` | Admin |

### Transactions
| Method | Endpoint | Access |
|--------|----------|--------|
| POST | `/api/transactions` | Cashier + Admin |
| GET | `/api/transactions` | All |

### Reports
| Method | Endpoint | Access |
|--------|----------|--------|
| GET | `/api/reports/sales` | Admin |
| GET | `/api/reports/top-products` | Admin |

---

## 🧪 Health Check

```
GET http://localhost:3000/health
```

Response:
```json
{
  "status": "ok",
  "message": "POS System is running!"
}
```

---

## 📄 License

MIT License — feel free to use this project for learning and portfolio purposes.