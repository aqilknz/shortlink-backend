# 🔗 ShortLink Backend API

RESTful API untuk aplikasi URL Shortener yang dibangun menggunakan Golang, PostgreSQL, dan Redis.

Aplikasi ini memungkinkan pengguna untuk:

- Register dan Login
- Membuat Short URL
- Menggunakan Custom Slug
- Mengelola URL yang dimiliki
- Tracking jumlah klik
- Logout dengan JWT Blacklist menggunakan Redis

---

## 🛠 Tech Stack

### Backend

- Golang
- Gin Framework
- PostgreSQL
- Redis
- JWT Authentication
- Swagger

### DevOps

- Docker
- Docker Compose

---

## 📂 Project Structure

```text
shortlink-backend/
│
├── cmd/
│   └── main.go
│
├── db/
│   ├── migrations/
│   └── seeds/
│
├── docs/
│
├── internal/
│   ├── config/
│   ├── controller/
│   ├── dto/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── router/
│   ├── service/
│   └── utils/
│
├── pkg/
│
├── docker-compose.yml
├── redis.conf
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

Pastikan telah terinstall:

- Git
- Go 1.23+
- Docker
- Docker Compose

Cek versi:

```bash
git --version
go version
docker --version
docker compose version
```

---

## 1. Clone Repository

```bash
git clone https://github.com/aqilknz/shortlink-backend.git

cd shortlink-backend
```

---

## 2. Install Dependencies

```bash
go mod tidy
```

---

## 3. Setup Environment Variables

Copy file environment:

```bash
cp .env.example .env
```

Sesuaikan konfigurasi jika diperlukan.

---

## 4. Start PostgreSQL & Redis

Project menggunakan Docker Compose untuk menjalankan:

- PostgreSQL
- Redis
- Migration
- Seeder

Jalankan:

```bash
docker compose up -d
```

Cek container:

```bash
docker ps
```

---

## 5. Run Application

```bash
go run cmd/main.go
```

Server akan berjalan pada:

```text
http://localhost:8080
```

---

## 📚 API Documentation

Swagger tersedia pada:

```text
http://localhost:8080/swagger/index.html
```

Generate ulang dokumentasi Swagger:

```bash
swag init -g cmd/main.go
```

---

## 📌 API Endpoints

### Authentication

| Method | Endpoint |
|----------|----------|
| POST | `/api/auth/register` |
| POST | `/api/auth/login` |
| DELETE | `/api/auth/logout` |

### Link Management

| Method | Endpoint |
|----------|----------|
| POST | `/api/links` |
| GET | `/api/links` |
| DELETE | `/api/links/:id` |
| GET | `/api/links/check-slug` |

### Redirect

| Method | Endpoint |
|----------|----------|
| GET | `/:slug` |

---

## Example Request

### Create Short Link

```http
POST /api/links
```

Header:

```text
Authorization: Bearer <token>
```

Body:

```json
{
  "original_url": "https://github.com",
  "custom_slug": "github"
}
```

Response:

```json
{
  "status": "success",
  "data": {
    "slug": "github",
    "short_url": "http://localhost:8080/github"
  }
}
```

---

## 🧠 Design Decisions

### Layered Architecture

Project dipisahkan menjadi beberapa layer:

```text
Controller
    ↓
Service
    ↓
Repository
    ↓
Database
```

Tujuannya agar kode lebih mudah dipelihara dan dikembangkan.

### Database Transaction

Database transaction digunakan pada proses yang membutuhkan konsistensi data sehingga tidak terjadi partial write ketika terjadi error.

### Redis Blacklist

Saat user logout, JWT akan disimpan ke Redis hingga masa berlaku token berakhir. Middleware akan memeriksa Redis sebelum mengizinkan akses endpoint yang dilindungi.

### Soft Delete

Link tidak langsung dihapus permanen sehingga histori data tetap dapat dipertahankan.

---

## 🧪 Running Tests

Menjalankan seluruh unit test:

```bash
go test ./...
```

Coverage:

```bash
go test ./... -cover
```

---

## 👨‍💻 Author

**Ahmad Aqil Khairun Nadzar**

GitHub: https://github.com/aqilknz

---

## 📄 License

This project is intended for educational and portfolio purposes.