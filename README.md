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
- `cmd/server` → Entry point  
- `internal/` → Core application logic  
- `k8s/` → Kubernetes configs  

---

## 🗄 Database
- MySQL used for persistence  
- Indexed fields for faster lookup  
- Tables:
  - `tasks (id, title, status)`

---

## 🐳 Run with Docker

```bash
docker-compose up --build

## 📬 Sample Request
Create Task
POST /task
{
  "title": "Learn Go",
  "status": "pending"
}
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending"
}

▶️ How to Run (Manual)
Clone the repo
Run:
docker-compose up --build
Open:
http://localhost:8080


🎯 Purpose

This project demonstrates:
REST API design
Layered architecture (handler → service → repository)
Database optimization basics
Containerized deployment using Docker
