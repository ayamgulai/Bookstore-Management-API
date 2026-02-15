
---

# 📚 Book & Category REST API

Project ini merupakan REST API sederhana untuk mengelola **kategori buku**, **data buku**, serta **autentikasi user** menggunakan HTTP protocol. API ini cocok digunakan sebagai backend untuk aplikasi manajemen buku atau pembelajaran REST API.

---

## 🚀 Kegunaan Project

API ini menyediakan fitur utama:

* Manajemen **Kategori Buku**
* Manajemen **Buku**
* **Login User** menggunakan username dan password
* Menampilkan relasi **buku berdasarkan kategori**

---

## ⚙️ Cara Menjalankan Project

1. Pastikan **Go** sudah terinstall
2. Clone repository ini
3. Masuk ke folder project
4. Jalankan aplikasi:

```bash
go run main.go
```

Secara default server akan berjalan di:

```
http://localhost:8080
```

---

## 🔐 Autentikasi User

### Login

**Endpoint**

```
POST /api/users/login
```

**Request Body**

```json
{
  "username": "admin",
  "password": "admin"
}
```

**Default Account**

* Username: `admin`
* Password: `admin`

> Endpoint ini digunakan untuk proses login user.

---

## 📂 API Categories

### 1️⃣ Get All Categories

Menampilkan seluruh kategori.

* **Endpoint**

```
GET /api/categories
```

---

### 2️⃣ Add Category

Menambahkan kategori baru.

* **Endpoint**

```
POST /api/categories
```

* **Request Body (contoh)**

```json
{
  "name": "Novel"
}
```

---

### 3️⃣ Get Category Detail

Menampilkan detail kategori berdasarkan ID.

* **Endpoint**

```
GET /api/categories/:id
```

---

### 4️⃣ Delete Category

Menghapus kategori berdasarkan ID.

* **Endpoint**

```
DELETE /api/categories/:id
```

---

### 5️⃣ Get Books by Category

Menampilkan daftar buku berdasarkan kategori tertentu.

* **Endpoint**

```
GET /api/categories/:id/books
```

---

## 📘 API Books

### 1️⃣ Get All Books

Menampilkan seluruh buku.

* **Endpoint**

```
GET /api/books
```

---

### 2️⃣ Add Book

Menambahkan buku baru.

* **Endpoint**

```
POST /api/books
```

* **Request Body**

```json
{
  "title": "Novel apa aja test",
  "description": "Novel inspiratif karya Andrea Hirata yang mengisahkan perjuangan anak-anak Belitung dalam meraih pendidikan.",
  "image_url": "https://example.com/images/laskar-pelangi.jpg",
  "release_year": 2024,
  "price": 75000,
  "total_page": 52,
  "category_id": 3
}
```

---

### 3️⃣ Get Book Detail

Menampilkan detail buku berdasarkan ID.

* **Endpoint**

```
GET /api/books/:id
```

---

### 4️⃣ Delete Book

Menghapus buku berdasarkan ID.

* **Endpoint**

```
DELETE /api/books/:id
```

---

## 📌 Catatan

* Pastikan `category_id` yang digunakan saat menambahkan buku **sudah ada**
* Semua response dikembalikan dalam format **JSON**
* API ini berjalan di port **8080**

---