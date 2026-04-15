# HNG Stage 1 - Personal API

A simple Go API deployed on a Linux VPS with Nginx reverse proxy.

## How to Run Locally

```bash
git clone https://github.com/your-username/hng-stage1.git
cd hng-stage1
go run main.go
```

## Endpoints

| Endpoint | Method | Response |
|----------|--------|----------|
| `/` | GET | `{"message": "API is running"}` |
| `/health` | GET | `{"message": "healthy"}` |
| `/me` | GET | `{"name": "...", "email": "...", "github": "..."}` |

## Live URL

https://hngstage0.chickenkiller.com
