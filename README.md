# 📝 Todo App Backend (Go + MongoDB + Fiber)

This is the backend service for a **Todo Application** built using:
- **Golang (Fiber v2)** → REST API
- **MongoDB Atlas** → Database
- **JWT + Cookies** → Authentication
- **Postman** → API Testing

---

## 🚀 Features
- User registration & login with secure JWT authentication
- Protected routes using middleware
- CRUD APIs for Todos
- MongoDB Atlas integration
- CORS enabled for Next.js frontend
- Graceful shutdown with context

---


---

## ⚙️ Setup & Installation

### 1. Clone Repository
```bash
git clone https://github.com/saurabhraut1212/nextjs_golang_project.git
cd nextjs_golang_project

```
---
## 2. Create .env file
- PORT=8000
- MONGO_URI=mongodb+srv://<user>:<password>@cluster0.xxxx.mongodb.net
- MONGO_DB=todo_db
- CORS_ORIGIN=http://localhost:3000
- JWT_SECRET=supersecretkey

---
## 3. Install Dependencies
```bash
go mod tidy
```

## 4. Run Server
```bash
go run cmd/server/main.go
```

---
## Testing APIs (Postman)
1. Import the Postman Environment & Collection:
   https://web.postman.co/workspace/My-Workspace~388302e8-5eb7-4c3f-821d-5523c39dad56/collection/26119400-3e56e119-d545-4170-a315-0cbf872dcfd2?action=share&source=copy-link&creator=26119400
2. Run APIs in this order:
   Register → Login → Me → Create Todo → List Todos → Update → Delete → Logout
   


