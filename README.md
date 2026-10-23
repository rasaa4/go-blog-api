# Go Blog API

A RESTful blog backend written in Go. Users can register, log in with JWT, and manage their own blog posts. Built with a clean, layered architecture and fully containerized with Docker Compose.

## Features

- User registration and login with **JWT authentication**
- Passwords hashed with **bcrypt** (never returned by the API)
- CRUD for blog posts; only the **author can update or delete** their post
- Request validation with **ozzo-validation**
- Structured errors with a custom `richerror` package and correct HTTP status codes
- **MySQL** with SQL migrations applied automatically on first start
- **Docker Compose** setup with a database health check
- Unit tests for the user use case

## Tech Stack

Go · chi router · MySQL 8 · JWT (golang-jwt) · bcrypt · ozzo-validation · Docker Compose

## Project Structure

```
cmd/                  # application entry point
config/               # environment configuration
db/                   # database connection
migrations/           # SQL migrations
internal/
  domain/             # core models
  dto/                # request/response objects + validation
  handler/            # HTTP handlers
  middleware/         # JWT authentication middleware
  usecase/            # business logic
  repository/         # data access (MySQL)
  richerror/          # structured error handling
  response/           # consistent JSON responses
```

## Getting Started

**Requirements:** Docker and Docker Compose

```bash
git clone https://github.com/rasaa4/go-blog-api.git
cd go-blog-api
cp .env.example .env      # then edit the values
docker compose up --build -d
```

The API runs on `http://localhost:8080`.

To stop it: `docker compose down` (add `-v` to also delete the database).

## API Endpoints

| Method | Endpoint      | Auth | Description                 |
|--------|---------------|------|-----------------------------|
| POST   | `/register`   | No   | Create a new user           |
| POST   | `/login`      | No   | Log in and receive a JWT    |
| POST   | `/posts`      | Yes  | Create a post               |
| GET    | `/posts`      | Yes  | List all posts              |
| GET    | `/posts/{id}` | Yes  | Get a post by ID            |
| PUT    | `/posts/{id}` | Yes  | Update a post (author only) |
| DELETE | `/posts/{id}` | Yes  | Delete a post (author only) |

Protected endpoints require the header: `Authorization: Bearer <token>`

## Example

```bash
# Register
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"rasa","password":"123456"}'

# Log in
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"rasa","password":"123456"}'

# Create a post (use the token from login)
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title":"My first post","content":"Hello from my Go blog API!"}'
```

Response format:

```json
{ "success": true, "data": { "id": 1, "title": "My first post", "content": "Hello from my Go blog API!", "user_id": 1 } }
```

## Running Tests

```bash
go test ./...
```