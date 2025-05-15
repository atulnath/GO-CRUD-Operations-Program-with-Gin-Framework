# GO-CRUD-Operations-Program-with-Gin-Framework

# Gin RESTful API Example

This project is a simple RESTful API server built with the [Gin](https://github.com/gin-gonic/gin) web framework in Go. It demonstrates how to handle different HTTP methods (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`) with query parameters and JSON payloads.

---

## 📦 Features

- `GET /ping` — Health check endpoint  
- `GET /me/:second_id/:first_id` — Returns `first_id` and `second_id` from URL params  
- `POST /userlogin` — Accepts user login data in JSON and returns it  
- `PUT /userlogin` — Updates complete user login data  
- `PATCH /userlogin` — Partially updates user login data  
- `DELETE /userlogin/:id` — Deletes a user by ID  

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (v1.16 or above recommended)

### Install Gin

```bash
go get -u github.com/gin-gonic/gin


PUT is used to replace the whole object.

PATCH is used to update a part of the object.

---
