# Toko API Fiber

REST API untuk manajemen produk, dibangun dengan Go + Fiber v3 + GORM.

## Tech Stack

- **Go** 1.26.2
- **Fiber v3** — HTTP framework
- **GORM** — ORM untuk database
- **MySQL** — Database
- **Google Wire** — Dependency Injection
- **Logrus** — Logging
- **Viper** — Configuration management
- **golang-migrate** — Database migration
- **go-playground/validator** — Request validation

## Cara Setup

### Prerequisites
- Go 1.26.2
- MySQL
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)

### 1. Clone repository
```bash
git clone https://github.com/Mpayy/toko-api-fiber.git
cd toko-api-fiber
```

### 2. Setup konfigurasi
```bash
cp config.env.example config.env
# Edit config.env sesuai konfigurasi database kamu
```

### 3. Jalankan migration
```bash
migrate -path db/migrations \
        -database "mysql://root:password@tcp(localhost:3306)/toko-api-fiber" up
```

### 4. Jalankan aplikasi
```bash
cd cmd/web/
go run .
```

---

## Endpoints

### Products

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/products` | Get semua produk |
| GET | `/api/products/:id` | Get produk by ID |
| POST | `/api/products` | Tambah produk baru |
| PUT | `/api/products/:id` | Update produk |
| PATCH | `/api/products/:id` | Partial Update produk |
| DELETE | `/api/products/:id` | Hapus produk |

> Semua endpoint membutuhkan header `X-API-Key: [nilai API key kamu]`

---

## Contoh Request & Response

### POST /api/products
**Request:**
```json
{
  "name": "product",
  "price": 10000,
  "stock": 10
}
```

**Response:**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "product",
    "price": 10000,
    "stock": 10,
    "created_at": "2026-06-15T02:27:55+07:00",
    "updated_at": "2026-06-15T02:27:55+07:00"
  }
}
```

### PATCH /api/products/:id
**Request:**
```json
{
  "name": "product baru",
  "price": 15000
}
```

**Response:**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "product baru",
    "price": 15000,
    "stock": 10,
    "created_at": "2026-06-15T02:27:55+07:00",
    "updated_at": "2026-06-15T02:36:55+07:00"
  }
}
```

### GET /api/products/:id — Not Found
**Response:**
```json
{
    "code": 404,
    "status": "Product with id = 1 not found"
}
```

### POST /api/products - Validation Failed
**Response:**
```json
{
    "code": 400,
    "status": "Bad Request",
    "errors": {
        "Name": "required",
        "Price": "required",
        "Stock": "required"
    }
}
```

### PATCH /api/products/:id - Validation Failed
**Response:**
```json
{
    "code": 400,
    "status": "Bad Request",
    "errors": {
        "Name": "min"
    }
}
```
---

## Project Structure

```
toko-api-fiber/
├── cmd/web/          # Entry point
├── db/migrations/    # SQL migration files
├── internal/
│   ├── config/       # Application config
│   ├── entity/       # GORM models
│   ├── model/        # Request & Response structs
│   ├── repository/   # Database layer
│   ├── usecase/      # Business logic
│   ├── delivery/http/# Controllers, Middleware & routing
│   └── util/         # Utilities
└── config.env.example
```