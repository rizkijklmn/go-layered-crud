# 📦 Go CRUD Layered Architecture

Project ini adalah implementasi sederhana **CRUD API** menggunakan **Golang** dengan pendekatan arsitektur **layered**:

### ✅ Arsitektur Folder
- `controller` – menangani HTTP request/response
- `service` – berisi logic / business process
- `repository` – mengakses dan mengelola data (ORM)
- `model` – struktur data dan relasi GORM

---

## 🔧 Tech Stack

| Library | Deskripsi |
|--------|-----------|
| [Gin](https://github.com/gin-gonic/gin) | HTTP Web Framework |
| [GORM](https://gorm.io/) | ORM (Object Relational Mapping) |
| [godotenv](https://github.com/joho/godotenv) | Mengelola konfigurasi via `.env` |
| [Air](https://github.com/air-verse/air) _(optional)_ | Hot reload saat development |

---

## 🚀 Installation Guide

### 1. Inisialisasi Module
```bash
go mod init {your_project_name}

# Gin - Web framework
go get github.com/gin-gonic/gin

# GORM - ORM
go get gorm.io/gorm
go get gorm.io/driver/mysql

# godotenv - untuk membaca file .env
go get github.com/joho/godotenv

# Install Air (Optional - untuk hot reload saat development)
go install github.com/air-verse/air@latest

