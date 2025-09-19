# Takehome CBI - Golang Backend API

A REST API backend built with Go and Gin framework, featuring JWT-based authentication, CRUD operations, input validation, and production-ready deployment configuration for Vercel.

## 📖 Overview

This project serves as the backend API for a complete authentication and item management system. It provides secure JWT authentication, comprehensive CRUD operations with validation, environment-based configuration, and seamless deployment to Vercel serverless functions.

## 🚀 Tech Stack

### Core Framework

-   **Go 1.25** - Modern, efficient programming language
-   **Gin Framework** - High-performance HTTP web framework
-   **Gin-Contrib/CORS** - Cross-Origin Resource Sharing middleware

### Authentication & Security

-   **golang-jwt/jwt/v5** - JSON Web Token implementation
-   **Environment Variables** - Secure credential management via godotenv
-   **Input Validation** - Built-in struct tag validation with go-playground/validator

### Development & Deployment

-   **godotenv** - Environment variable loading for development
-   **Vercel** - Serverless deployment platform
-   **CORS Configuration** - Frontend integration support

## ✨ Features

### 🔑 Authentication System

-   **JWT-based Authentication** - Secure token-based user authentication
-   **Environment-based Credentials** - Username/password stored in environment variables
-   **Token Validation Middleware** - Automatic route protection
-   **Bearer Token Support** - Standard Authorization header handling

### 📝 Item Management (CRUD)

-   **Get All Items** - Retrieve complete item list with structured responses
-   **Get Single Item** - Fetch specific item by ID
-   **Create Item** - Add new items with validation
-   **Update Item** - Modify existing items with validation
-   **Delete Item** - Remove items with confirmation

### 🛡️ Input Validation

-   **Price Validation** - Ensures positive values (gt=0)
-   **Name Validation** - Required field with length constraints (1-100 characters)
-   **ID Validation** - Non-negative integer validation (gte=0)

### 🌐 Production Ready

-   **Vercel Deployment** - Serverless function configuration
-   **CORS Support** - Cross-origin request handling
-   **Environment Configuration** - Development and production environment support
-   **Health Check Endpoint** - API status monitoring

## 🚦 Getting Started

### Prerequisites

-   Go 1.25 or higher
-   Git for version control

### Installation

1. **Clone the repository**

    ```bash
    git clone https://github.com/alifrachmat2002/takehome-cbi-golang-be.git
    cd takehome-cbi-golang-be
    ```

2. **Install dependencies**

    ```bash
    go mod download
    ```

3. **Environment Setup**
   Create a `.env` file based on `.env.example`:

    ```env
    AUTH_USERNAME=your_admin_username
    AUTH_PASSWORD=your_secure_password
    JWT_SECRET=your_jwt_secret_key_minimum_32_characters
    PORT=8080
    ```

4. **Run the development server**

    ```bash
    go run main.go
    ```

5. **API is ready**
   Server running on [http://localhost:8080](http://localhost:8080)

## � API Endpoints

### Authentication

| Method | Endpoint | Description         | Auth Required |
| ------ | -------- | ------------------- | ------------- |
| POST   | `/login` | User authentication | ❌            |

### Items Management

| Method | Endpoint         | Description          | Auth Required |
| ------ | ---------------- | -------------------- | ------------- |
| GET    | `/api/items`     | Get all items        | ✅            |
| GET    | `/api/items/:id` | Get specific item    | ✅            |
| POST   | `/api/items`     | Create new item      | ✅            |
| PUT    | `/api/items/:id` | Update existing item | ✅            |
| DELETE | `/api/items/:id` | Delete item          | ✅            |

### Health Check

| Method | Endpoint | Description       | Auth Required |
| ------ | -------- | ----------------- | ------------- |
| GET    | `/`      | API health status | ❌            |

## 📝 API Usage Examples

### Authentication

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "your_username",
    "password": "your_password"
  }'
```

**Response:**

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Create Item

```bash
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Sample Item",
    "price": 29999
  }'
```

**Response:**

```json
{
    "status": 200,
    "message": "Data Created successfully!",
    "data": {
        "id": 3,
        "name": "Sample Item",
        "price": 29999
    }
}
```

## 🚀 Deployment

### Vercel Deployment

This project is configured for seamless deployment to Vercel:

1. **Install Vercel CLI**

    ```bash
    npm i -g vercel
    ```

2. **Deploy to Vercel**

    ```bash
    vercel --prod
    ```

3. **Configure Environment Variables**
   Set these in your Vercel project dashboard:
    - `AUTH_USERNAME`
    - `AUTH_PASSWORD`
    - `JWT_SECRET`
    - `GIN_MODE=release`


## 🔒 Security Features

-   **Environment-based Authentication** - No hardcoded credentials
-   **JWT Token Validation** - Secure API access control
-   **Input Validation** - Protection against invalid data
-   **CORS Configuration** - Controlled cross-origin access
-   **Production Mode** - Optimized for production deployment

## 🔗 Connect

-   **Portfolio**: [Alif Rachmat Illahi](https://portofolio-web-phi-wine.vercel.app/)
-   **LinkedIn**: [Alif Rachmat Illahi](https://www.linkedin.com/in/alifrachmat/)

---
