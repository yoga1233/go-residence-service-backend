# Backend Aplikasi Perumahan

Backend untuk aplikasi perumahan yang menyediakan fitur manajemen berita, layanan (seperti AC dan tukang kebun), serta fitur laporan (seperti laporan pohon tumbang dan status penanganannya). Dibangun menggunakan **Golang**, **Fiber**, **GORM**, **JWT**, dan **BCrypt**.

## Fitur

- **Autentikasi Pengguna**: Login dan registrasi dengan enkripsi password menggunakan BCrypt.
- **JWT Authentication**: Menggunakan JWT untuk mengamankan API.
- **Manajemen Berita**: Tambah, edit, dan hapus berita perumahan.
- **Manajemen Layanan**: Daftar layanan perumahan yang tersedia (seperti AC dan tukang kebun).
- **Laporan Warga**: Fitur untuk membuat laporan masalah (misalnya pohon tumbang).
- **Status Penanganan Laporan**: Melacak status laporan yang dibuat oleh warga.

---

## Teknologi yang Digunakan

- [Golang](https://golang.org/) - Bahasa pemrograman utama
- [Fiber](https://gofiber.io/) - Web framework untuk Golang
- [GORM](https://gorm.io/) - ORM untuk database
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token untuk autentikasi
- [BCrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Enkripsi password

---




