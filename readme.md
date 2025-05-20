# Go CRUD Layered Architecture

Project ini adalah implementasi sederhana CRUD API menggunakan **Golang** dengan pendekatan arsitektur layered:

- `controller`
- `service`
- `repository`
- `model`

Menggunakan library:
- [`Gin`](https://github.com/gin-gonic/gin) — HTTP web framework
- [`GORM`](https://gorm.io/) — ORM untuk Golang
- [`godotenv`](https://github.com/joho/godotenv) — untuk mengelola environment variables

Optional
- [`Air`](https://github.com/air-verse/air) — untuk hot reload

# Installation Guide
- Inisialisasi module (jika belum)
go mod init go-crud

- Install Gin (Web Framework)
go get github.com/gin-gonic/gin

- Install GORM (ORM Golang)
go get gorm.io/gorm
go get gorm.io/driver/mysql

- Install godotenv (untuk membaca file .env)
go get github.com/joho/godotenv

Optional
- install Air untuk hot reload
go install github.com/air-verse/air@latest

