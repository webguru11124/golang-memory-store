
# In-Memory Data Store API Documentation

## Overview
This API provides functionalities similar to Redis for storing key-value pairs and lists with support for TTL (Time-To-Live), Authentication (JWT), and Data Persistence.

### Base URL
```
http://localhost:8080
```

## Authentication

### Generate Token
```
POST /token
```
- **Request Body:** 
```json
{
    "username": "testuser"
}
```
- **Response:**
```json
{
    "token": "YOUR_JWT_TOKEN"
}
```

---

## String Operations

### Set Key-Value Pair (With TTL)
```
POST /set
```
- **Headers:** `Authorization: Bearer YOUR_JWT_TOKEN`
- **Request Body:**
```json
{
    "key": "foo",
    "value": "bar",
    "ttl": 60
}
```
- **Response:**
```json
{
    "success": true,
    "message": "Set Successfully"
}
```

---

### Get Key
```
GET /get/{key}
```
- **Headers:** `Authorization: Bearer YOUR_JWT_TOKEN`
- **Example Request:**
```
GET /get/foo
```
- **Response:**
```json
{
    "value": "bar"
}
```

---

### Delete Key
```
DELETE /delete/{key}
```
- **Headers:** `Authorization: Bearer YOUR_JWT_TOKEN`
- **Example Request:**
```
DELETE /delete/foo
```
- **Response:**
```json
{
    "success": true,
    "message": "Deleted Successfully"
}
```

---

## List Operations

### Push to List
```
POST /list/push
```
- **Headers:** `Authorization: Bearer YOUR_JWT_TOKEN`
- **Request Body:**
```json
{
    "key": "mylist",
    "value": "item1"
}
```
- **Response:**
```json
{
    "success": true,
    "message": "Item Pushed Successfully"
}
```

---

### Pop from List
```
POST /list/pop
```
- **Headers:** `Authorization: Bearer YOUR_JWT_TOKEN`
- **Request Body:**
```json
{
    "key": "mylist"
}
```
- **Response:**
```json
{
    "value": "item1"
}
```

---

## Data Persistence
- Data is saved to a file at regular intervals or during shutdown.
- The file-based persistence feature allows data restoration on server restart.

## Note
- All API requests must include a valid JWT token in the `Authorization` header.
- This API is implemented in Go with modular architecture for better maintainability and scalability.
