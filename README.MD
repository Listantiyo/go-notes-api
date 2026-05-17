# 📝 Notes API (Base Project untuk Praktik CI/CD)

Proyek ini adalah REST API sederhana berbasis Go (Golang) yang mengelola catatan (*notes*). Proyek ini sengaja dirancang dengan struktur minimalis dan fitur esensial sebagai **proyek dasar (base project) untuk kebutuhan praktik implementasi CI/CD menggunakan GitHub Actions**.

## 🚀 Fitur Utama & Struktur Data
* **Framework Router:** `go-chi/chi` (dengan penanganan routing dinamis).
* **Database:** SQLite (disimpan lokal dalam bentuk file biner di dalam folder `./database`).
* **ID Generasi:** UUID V4 (menggunakan paket `://github.com`).

### Skema Tabel (`notes`)

| Kolom | Tipe Data | Keterangan |
| :--- | :--- | :--- |
| `id` | `TEXT` | Primary Key (UUID V4 String) |
| `title` | `TEXT` | Judul Catatan |
| `descriptions` | `TEXT` | Isi / Deskripsi Catatan |

---

## 🛣️ Endpoint API

Aplikasi ini menyediakan dua endpoint utama yang menangani data secara dinamis:

### 1. Ambil Data Catatan (`GET`)
Mengambil semua data catatan atau spesifik berdasarkan ID menggunakan Query Parameter opsional.
* **URL:** `/note` atau `/note?id={uuid}`
* **Metode:** `GET`

### 2. Buat Catatan Baru (`POST`)
Memasukkan catatan baru ke dalam database SQLite. ID akan digenerasikan otomatis oleh server menggunakan UUID V4.
* **URL:** `/note`
* **Metode:** `POST`
* **Format Body (JSON):**
  ```json
  {
    "title": "Judul Catatan Anda",
    "descriptions": "Isi deskripsi lengkap catatan."
  }
  ```

---

## 🛠️ Cara Menjalankan Proyek Secara Lokal

### Menggunakan Go Lokal
1. Pastikan dependensi diselaraskan:
   ```bash
   go mod tidy
   ```
2. Jalankan server utama:
   ```bash
   go run ./cmd/server
   ```
3. Aplikasi akan berjalan di port `http://localhost:8080`.

### Menggunakan Docker Compose (Direkomendasikan untuk CI/CD)
Proyek ini sudah dilengkapi dengan konfigurasi volume agar file database SQLite tidak hilang saat container dimatikan:
```bash
# Build image lokal dan jalankan container di latar belakang
docker compose up -d --build
```
*File database akan otomatis tersinkronisasi di komputer host Anda pada folder `./data_lokal`.*

---

## 🧪 Panduan Pengujian API via PowerShell

Anda dapat menguji fungsionalitas kedua endpoint langsung melalui terminal Windows PowerShell menggunakan perintah `Invoke-RestMethod` berikut:

### Tes POST (Membuat Catatan Baru)
*Pastikan parameter `-Body` diberi jarak spasi dan dibungkus tanda kutip tunggal di luar.*
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/note" -Method Post -ContentType "application/json" -Body '{"title": "Jangan Tidur", "descriptions":"Jika Anda tidur, server akan mati."}'
```

### Tes GET (Mengambil Semua Catatan)
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/note" -Method Get
```

### Tes GET Spesifik ID (Menggunakan Query Parameter)
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/note?id=PASTE_UUID_DI_SINI" -Method Get
```

---

## 🔄 Target Praktik CI/CD (GitHub Actions)
Repository ini siap dipasangkan dengan alur kerja otomatisasi (*workflow*) GitHub Actions untuk melakukan:
1. **Linter & Code Checking:** Validasi format kode Go menggunakan `golangci-lint`.
2. **Automated Testing:** Menjalankan unit testing jika tersedia.
3. **Docker Build Integration:** Otomatisasi proses pengujian *build* `Dockerfile` multi-stage untuk memastikan aplikasi selalu siap dideploy tanpa kendala *compile*.
