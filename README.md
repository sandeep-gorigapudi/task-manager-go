# Task Manager API (Go + Gin)

## 🚀 Overview
A backend REST API for managing tasks, built using Go (Gin), MySQL, and Docker.

---

## ⚙️ Features
- Create Task
- Get All Tasks
- Update Task
- Delete Task

---

## 🧱 Architecture
Client → Gin Router → Handler → Service → Repository → MySQL


---

## 🛠 Tech Stack
- Go (Golang)
- Gin Framework
- MySQL
- Docker
- Docker Compose

---

## 📂 Project Structure
cmd/server → Entry point
internal/ → Core application logic
k8s/ → Kubernetes configs

## 🐳 Run with Docker

```bash
docker-compose up --build


| Method | Endpoint  | Description   |
| ------ | --------- | ------------- |
| POST   | /task     | Create task   |
| GET    | /tasks    | Get all tasks |
| PUT    | /task/:id | Update task   |
| DELETE | /task/:id | Delete task   |
