# Usage Guide

## Start the Application

Run the server using

```bash
go run main.go
```

Server will start on:

```text
http://localhost:8080
```

---

# API Endpoints

| Method | Endpoint      | Description          |
|--------|---------------|----------------------|
| POST   | /todos        | Create a todo        |
| GET    | /todos        | List todos           |
| GET    | /todos/{id}   | Get todo by ID       |
| PUT    | /todos/{id}   | Update todo          |
| DELETE | /todos/{id}   | Delete todo          |

---

# 1. Create Task

## Request

```bash
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{
  "task":"Learn Go",
  "due_date":"2026-05-20"
}'
```

## Response

```json
{
  "id": 1,
  "task": "Learn Go",
  "due_date": "2026-05-20T00:00:00Z",
  "completed": false
}
```

---

# 2. Get Todo by ID

## Request

```bash
curl http://localhost:8080/todos/1
```

## Response

```json
{
  "id": 1,
  "task": "Learn Go",
  "due_date": "2026-05-20T00:00:00Z",
  "completed": false
}
```


---

# 3. List Todos

Completed tasks are excluded by default.

## Request

```bash
curl http://localhost:8080/todos
```

## Response

```json
[
  {
    "id": 1,
    "task": "Learn Go",
    "due_date": "2026-05-20T00:00:00Z",
    "completed": false
  }
]
```

---

# Include Completed Todos

## Request

```bash
curl "http://localhost:8080/todos?include_completed=true"
```

---

# 4. Update Todo

## Request

```bash
curl -X PUT http://localhost:8080/todos/1 \
-H "Content-Type: application/json" \
-d '{
  "task":"Practice Go",
  "due_date":"2026-05-25",
  "completed":true
}'
```

## Response

```json
{
  "id": 1,
  "task": "Practice Go",
  "due_date": "2026-05-25T00:00:00Z",
  "completed": true
}
```

---

# 5. Delete Todo

## Request

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Response

```text
204 No Content
```

---

# Notes

- Due date format must be:
  
text
YYYY-MM-DD


Example:

text
2026-05-20


- Todos are sorted by due date in ascending order.
- Data is stored in-memory and will be lost after server restart.