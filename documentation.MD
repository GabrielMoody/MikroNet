# API Endpoints

## Authentication

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

Response Body
```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImdhYnJpZWxAbWFpbC5jb20iLCJleHAiOjE3MzU2MjczOTAsImlkIjoiZDQzZDBlMDktNmE3ZS00NGYxLWEwYjMtZGYzMTI3ZmZiN2U3IiwiaXNzIjoiIiwicm9sZSI6InVzZXIifQ.WfcxHBaesi5NT4ObTOhsZcU_lNemo0Pu7VG5xvKaPVw"
  },
  "status": "success"
}
```

2. User Registration

```http request
POST /api/auth/register/user
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

4. Request Reset Password
```http request
POST /api/auth/reset-password
Content-Type: application/json
```

Request Body
```json
{
  "email": "contoh@gmail.com"
}
```

5. Reset Password
```http request
PUT /api/auth/reset-password/{token}
Content-Type: application/json
```

Request Body
```json
{
  "password": "*********",
  "password_confirmation": "*********"
}
```

6. Change Password
```http request
PUT /api/auth/change-password
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
  "old_password": "*********",
  "new_password": "*********",
  "new_password_confirmation": "*********"
}
```

## User

---

7. Get User Data
```http request
GET /api/user/
Content-Type: application/json
Authorization: Bearer <token>
```

8. Update User Data
```http request
PUT /api/user/
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
  "age": 54
}
```

9. Delete User
```http request
PATCH /api/user/
Content-Type: application/json
Authorization: Bearer <token>
```

10. Review Driver
```http request
POST /api/user/review/{driver_id}
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
  "rating": 5,
  "comment": "good driver"
}
```

## Driver

---

11. Get Driver Data
```http request
GET /api/driver/
Content-Type: application/json
Authorization: Bearer <token>
```

12. Update Driver Data
```http request
PUT /api/driver/
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
    "first_name": "first name",
    "last_name": "last name",
    "email": "contoh@gmail.com",
    "phone_number": "089279213091",
    "date_of_birth": "2000-01-01",
    "age": 53,
    "gender": "Male"
}
```

13. Get status driver
```http request
GET /api/driver/status
Content-Type: application/json
Authorization: Bearer <token>
```

14. Update status driver
```http request
PUT /api/driver/status
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
    "status": "on"
}
```

15. Get available seats
```http request
GET /api/driver/seats
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
    "status": "success",
    "data": {
        "seats": 3
    }
}
```

16. Update available seats
```http request
PUT /api/driver/seats
Content-Type: application/json
Authorization: Bearer <token>
```

Request Body
```json
{
    "seats": 3
}
```

Response Body

```json
{
    "status": "success",
    "data": {
        "seats": 3
    }
}
```

## Admin

---

17. Get All Users
```http request
GET /api/dashboard/users
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "count": 1,
    "users": [
      {
        "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "first_name": "first name",
        "last_name": "last name",
        "email": "contoh@gmail.com"
      }
    ]
  }
}
```

18. Get All Drivers
```http request
GET /api/dashboard/drivers
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "count": 1,
    "drivers": [
      {
        "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "first_name": "first name",
        "last_name": "last name",
        "email": "contoh@gmail.com"
      }
    ]
  }
}
```

19. Get Users details
```http request
GET /api/dashboard/users/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
    "first_name": "first name",
    "last_name": "last name",
    "email": "contoh@gmail.com"
  }
}
```

20. Get Drivers details
```http request
GET /api/dashboard/drivers/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
    "first_name": "first name",
    "last_name": "last name",
    "email": "contoh@gmail.com"
  }
}
```

21. Get All Owners
```http request
GET /api/dashboard/owners
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "count": 1,
    "owners": [
      {
        "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "first_name": "first name",
        "last_name": "last name",
        "email": "contoh@gmail.com"
      }
    ]
  }
}
```

22. Get Blocked Accounts
```http request
GET /api/dashboard/blocked
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "count": 1,
    "blocked": [
      {
        "id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "first_name": "first name",
        "last_name": "last name",
        "email": "contoh@gmail.com"
      }
    ]
  }
}
```

23. Block Account
```http request
POST /api/dashboard/block/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

24. Get All Reviews
```http request
GET /api/dashboard/reviews
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "count": 1,
    "reviews": [
      {
        "id": "1",
        "driver_id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "user_id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
        "rating": 5,
        "comment": "good driver"
      }
    ]
  }
}
```

25. Get Review By ID
```http request
GET /api/dashboard/reviews/{id}
Content-Type: application/json
Authorization: Bearer <token>
```

Response Body

```json
{
  "status": "success",
  "data": {
    "id": "1",
    "driver_id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
    "user_id": "d43d0e09-6a7e-44f1-a0b3-df3127ffb7e7",
    "rating": 5,
    "comment": "good driver"
  }
}
```