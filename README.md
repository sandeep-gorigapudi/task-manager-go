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

## 📬 Sample Request

### Create Task

POST /task

```json
{
  "title": "Learn Go",
  "status": "pending"
}
Response
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending"
}


---

## 🔹 Run Instructions

```markdown
## ▶️ How to Run

1. Clone the repo
2. Run:

```bash
docker-compose up --build

Server runs on:
http://localhost:8080


---

## 🔹 Database

```markdown
## 🗄 Database

- MySQL used for persistence
- Indexed fields for faster lookup
- Tables:
  - tasks (id, title, status)

## 🎯 Purpose

This project demonstrates:
- REST API design
- Layered architecture (handler → service → repo)
- Database optimization basics
- Containerized deployment using Docker
