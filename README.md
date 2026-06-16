# Toko API Fiber

REST API sederhana untuk manajemen produk dan autentikasi user, dibangun menggunakan Go, Fiber v3, dan GORM dengan implementasi Clean Architecture. Proyek ini dibuat sebagai portofolio untuk mempraktikkan backend development menggunakan ekosistem Go modern.

## 🚀 Fitur Utama

- **Clean Architecture** (Pemisahan layer Controller, Usecase, dan Repository)
- **Dependency Injection** menggunakan Google Wire
- **Autentikasi JWT** (JSON Web Token)
- **Manajemen User** (Register, Login, Get Current User, Logout)
- **Manajemen Produk** (CRUD lengkap termasuk Partial Update / PATCH)
- **Validasi Input** dengan `go-playground/validator` (termasuk custom rule validasi kelipatan harga)
- **Middleware Terpusat** untuk Logging dan penanganan Error

## 🛠️ Tech Stack

- **Go** 1.26
- **Fiber v3** — HTTP framework
- **GORM** — ORM untuk database
- **MySQL** — Database
- **Google Wire** — Dependency Injection
- **golang-jwt** — Authentication
- **Logrus** — Logging
- **Viper** — Configuration management
- **golang-migrate** — Database migration
- **go-playground/validator** — Request validation

## ⚙️ Cara Setup

### Prerequisites

- Go 1.26+
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
# Edit config.env sesuai konfigurasi database dan secret key JWT kamu
```

### 3. Jalankan migration

```bash
migrate -path db/migration \
        -database "mysql://username:password@tcp(localhost:3306)/toko-api-fiber" up
```

### 4. Generate Wire (Optional jika ada perubahan dependency)

```bash
cd cmd/web/
wire
```

### 5. Jalankan aplikasi

```bash
cd cmd/web/
go run .
```

---

## 📡 Endpoints

### 👤 Users

| Method | Endpoint              | Deskripsi                               | Auth |
| ------ | --------------------- | --------------------------------------- | ---- |
| POST   | `/api/users/register` | Mendaftarkan user baru                  | ❌   |
| POST   | `/api/users/login`    | Login dan mendapatkan JWT token         | ❌   |
| GET    | `/api/users/current`  | Mendapatkan data user yang sedang login | ✅   |
| DELETE | `/api/users/logout`   | Logout (menghapus token)                | ✅   |

### 📦 Products

| Method | Endpoint            | Deskripsi             | Auth |
| ------ | ------------------- | --------------------- | ---- |
| GET    | `/api/products`     | Get semua produk      | ❌   |
| GET    | `/api/products/:id` | Get produk by ID      | ❌   |
| POST   | `/api/products`     | Tambah produk baru    | ✅   |
| PUT    | `/api/products/:id` | Update produk         | ✅   |
| PATCH  | `/api/products/:id` | Partial Update produk | ✅   |
| DELETE | `/api/products/:id` | Hapus produk          | ✅   |

> Endpoint dengan Auth ✅ membutuhkan header `Authorization: Bearer <token_jwt>`

---

## 📄 Contoh Request & Response

### Register User (`POST /api/users/register`)

**Request:**

```json
{
  "username": "johndoe",
  "email": "johndoe@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "data": {
    "id": 1,
    "username": "johndoe",
    "created_at": "2026-06-16T10:00:00Z",
    "updated_at": "2026-06-16T10:00:00Z"
  }
}
```

### Login (`POST /api/users/login`)

**Response:**

```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Get All Products (`GET /api/products?page=2&size=10`)

**Response:**

```json
{
  "data": [
    {
      "id": 20,
      "name": "Laptop Gaming",
      "price": 15000000,
      "stock": 10,
      "created_at": "2026-06-16T10:00:00Z",
      "updated_at": "2026-06-16T10:00:00Z"
    }
  ],
  "paging": {
    "page": 2,
    "total_pages": 2,
    "total_items": 20
  }
}
```

### Validation Error Response

**Response (Status 400):**

```json
{
  "errors": {
    "email": "must be a valid email",
    "password": "must be at least 8 characters long"
  }
}
```

### Unauthorized Response

**Response (Status 401):**

```json
{
  "errors": "Unauthorized"
}
```

---

## 📂 Project Structure

```text
toko-api-fiber/
├── cmd/web/          # Entry point aplikasi & Wire injector
├── db/migration/     # File migrasi SQL
├── internal/
│   ├── config/       # Konfigurasi aplikasi (Database, Fiber, Viper, dsb)
│   ├── delivery/     # Layer HTTP (Controllers, Middleware, Routing)
│   ├── entity/       # Struktur representasi tabel database (GORM)
│   ├── exception/    # Custom error handling & custom validation
│   ├── model/        # Struktur DTO (Request/Response)
│   ├── repository/   # Operasi database
│   ├── usecase/      # Business logic / Service layer
│   └── util/         # Utility functions (contoh: JWT parser)
├── config.env.example# Template environment variables
└── README.md
```
