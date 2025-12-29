# 📒 Address Book Backend (Go + Gin + Gorm + PostgreSQL)

A practice project to explore Go backend development: An Address Book backend system built with **Go (Golang)**, following clean architecture principles, JWT auth, CSV export, asynchronous processing, search, filters, pagination, and real email sending.

---

## Tech Stack

| Layer | Technology |
|-------|-------------|
| Language | Go (Golang) |
| Web Framework | Gin |
| Database | PostgreSQL |
| ORM | GORM |
| Authentication | JWT |
| Email Service | SMTP (Gmail / App Password) |
| CSV Export | Encoding/CSV, async goroutines |
| Architecture | Layered: Controller → Service → Repository |

---

## 📂 Project Structure

```

address-book-backend/
├── cmd/server/main.go
├── internal/
│   ├── config/
│   ├── db/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   ├── controllers/
│   ├── middleware/
│   ├── router/
│   ├── utils/
├── exports/            # generated CSV files
├── .env
└── go.mod

```

---

## 🔐 Authentication

| Endpoint | Method | Description |
|----------|---------|--------------|
| `/auth/register` | POST | Register a new user |
| `/auth/login` | POST | Login & receive JWT token |

Headers required:
```

Authorization: Bearer <token>

```

---

## 📒 Address Management (Protected Endpoints)

| Endpoint | Method | Description |
|----------|---------|--------------|
| `/addresses` | GET | Get address data |
| `/addresses` | POST | Create new address |
| `/addresses/:id` | PUT | Update address |
| `/addresses/:id` | DELETE | Soft delete |

**Query Params for Search & Filter**
```

GET /addresses/filter?search=yash&city=Surat&page=2&limit=5

```

---

## 📤 CSV Export Feature

### Custom Export Request
```json

{
  "fields": ["first_name", "email", "city"],
  "send_to": "manager@example.com"
}

````

### 📌 Flow

* Async goroutine
* Generate CSV file
* Email with attachment

### Public Download

```

GET /downloads/filename.csv

```

---

## Run Locally

### 1️⃣ Clone

```bash

git clone https://github.com/yash-sojitra-20/address-book-backend.git
cd address-book-backend

```

### 2️⃣ Setup `.env`

```env

DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASS=yourpassword
DB_NAME=yourdbname

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=app-password

APP_URL=http://localhost:8080

```

### 3️⃣ Run

```bash

go run cmd/server/main.go

```
