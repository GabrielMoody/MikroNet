# MikroNet Endpoints

## Authentication

1. User registration
```http
POST /api/registration
```

Request Body
```json
{
  "nama_lengkap": "Gabriel Moody Waworundeng",
  "email": "gabriel@gmail.com",
  "nomor_telepon": "+6282201923928",
  "kata_sandi": "********",
  "konfirmasi_kata_sandi": "********"
}
```

2. User Login

```http
POST /api/login
```

Request Body
```json
{
  "email": "gabriel@gmail.com",
  "kata_sandi": "********"
}
```

# Mobile

3. Mencari Mikrolet

4. Pesan Mikrolet

5. Carter Mikrolet

6. Riwayat Perjalanan

```http
Authorizations: Bearer <token> 
GET /api/{user}/histories
```

Response Body

```json
{
  "status": "selesai",
  "tanggal": "01-01-2024",
  "harga": 6000,
  "lokasi_awal": "Megamall",
  "lokasi_akhir": "Kost Kleak"
}
```

7. Profil

```http
Authorizations: Bearer <token> 
GET /api/{user}
```

Response Body
```json
{
  "nama_lengkap": "Gabriel Moody Waworundeng",
  "email": "gabriel@gmail.com",
  "nomor_telepon": "+6282201923928",
  "jenis_kelamin": "Laki-Laki",
  "tanggal_lahir": "01-01-2000",
  "gambar_url": "https://public-cloud-image.com/user-image.jpg"
}
```

8. Edit profile

```http
Authorizations: Bearer <token> 
PATCH /api/edit-profile/{user}
```

Request Body
```json
{
  "nama_lengkap": "Gabriel Moody Waworundeng",
  "email": "gabriel@gmail.com",
  "nomor_telepon": "+6282201923928",
  "jenis_kelamin": "Laki-Laki",
  "tanggal_lahir": "01-01-2000",
  "gambar_url": "https://public-cloud-image.com/user-image.jpg"
}
```

# Web