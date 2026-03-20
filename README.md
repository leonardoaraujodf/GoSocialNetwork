# GoSocialNetwork

A Go-based social network API built as a learning project.

## Features

- User authentication (JWT) and user management
- Follow/unfollow functionality
- Posts with comments
- Rate limiting
- Redis caching (toggleable via `REDIS_ENABLED`)
- Email notifications (SendGrid)

## Tech Stack

Go, PostgreSQL, Redis, chi router, zap logging, Swagger docs

## Quick Start

```bash
# Run migrations
make migrate-up

# Seed database
make seed

# Start server
air
```

## API Docs

Swagger UI available at `http://localhost:8080/v1/swagger/`

## Testing

```bash
make test
```

## Benchmarking

```bash
REDIS_ENABLED=true npx autocannon http://localhost:8080/v1/users/1 \
  --connections 50 --duration 5 \
  -H "Authorization: Bearer <token>"
```
