# API Endpoints

---
1. Login

```http request
GET /api/auth/login
Content-Type: application/json
```

Request Body
```json
{
  "email": "gabriel@gmail.com",
  "password": "**********"
}
```

2. User Registration

```http request
POST /api/auth/register
Content-Type: application/json
```

Request Body
```json
{
  "first_name": "first name",
  "last_name": "last name",
  "email": "mustbeanemail@mail.com",
  "phone_number": "089279213091",
  "date_of_birth": "2000-01-01",
  "age": 54,
  "password": "*********",
  "password_confirmation": "*********"
}
```

3. Driver Registration

```http request
POST /api/auth/register/driver
Content-Type: application/json
```

Request Body
```json
{
  "first_name": "first name",
  "last_name": "last name",
  "email": "mustbeanemail@mail.com",
  "phone_number": "089279213091",
  "date_of_birth": "2000-01-01",
  "age": 54,
  "password": "*********",
  "password_confirmation": "*********"
}
```

4. Get User Data
```http request
GET /api/profile/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

5. Update User Data
```http request
PATCH /api/profile/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
  "first_name": "first name",
  "last_name": "last name",
  "email": "mustbeanemail@mail.com",
  "gender": "male",
  "date_of_birth": "2000-01-01",
  "age": 54,
  "image": "imagefile.jpg"
}
```

6. Delete User
```http request
PATCH /api/profile/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

7. 