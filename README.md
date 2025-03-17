# Golang Devbook API

This repository contains the step-by-step implementation of a **Golang API** using **Gorilla Mux** for routing.

## 🔹 Lessons

<details>
  <summary><strong>📌 Lesson 69: packages, main and routes</strong></summary>

### 1️⃣ Comands:
```sh
go mod init api
touch main.go
code main.go
go get github.com/gorilla/mux
mkdir src
cd src
mkdir router && cd router && touch router.go && code router.go
```
</details>

<details>
  <summary><strong>📌 Lesson 70: packages, routes and controllers</strong></summary>

### 1️⃣ Comands:
```sh
mkdir -p src/router/routes && touch src/router/routes/routes.go
code src/router/routes/routes.go
touch src/router/routes/user.go
code src/router/routes/user.go
mkdir -p src/controllers
touch src/controllers/usuarios.go
code src/controllers/usuarios.go

```
</details>

<details>
  <summary><strong>📌 Lesson 71: Inseting routes inside the router</strong></summary>
  Play the routes into the router

</details>

<details>
  <summary><strong>📌 Lesson 72: Testing routes in postman</strong></summary>

  ### 1️⃣ Comands:
```sh
mkdir postman
```
</details>

<details>
  <summary><strong>📌 Lesson 73: Model user and init database</strong></summary>

  ### 1️⃣ Comands:
```sh
mkdir -p src/models && touch src/models/user.go && code src/models/user.go
mkdir -p src/sql && touch src/sql/sql.sql && code src/sql/sql.sql
mysql -u golang -p
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL UNIQUE,
    created_in TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;
desc users;
```
</details>

<details>
  <summary><strong>📌 Lesson 74: Package config</strong></summary>

  ### 1️⃣ Comands:
```sh
mkdir -p src/config && touch src/config/config.go && code src/config/config.go
touch .env && code .env
go get github.com/joho/godotenv
touch .env && code .env
touch .gitignore && code .gitignore
touch .env.example && code .env.example
```
</details>

