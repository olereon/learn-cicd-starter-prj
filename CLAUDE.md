# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Notely is a note-taking web application built with Go for the "Learn CICD" course on Boot.dev. It features user authentication, note creation/retrieval, and a simple web interface.

## Common Development Commands

### Build and Run
```bash
# Local development (non-database mode)
go build -o notely && ./notely

# Production build (Linux AMD64)
./scripts/buildprod.sh
```

### Database Code Generation
```bash
# Generate type-safe database code from SQL queries
sqlc generate
```

### Environment Setup
Create a `.env` file with:
```
PORT="8080"
DATABASE_URL="<optional-turso-db-url>"
```

## Architecture

### Backend Structure
- **Entry Point**: `main.go` - Sets up Chi router, middleware, and API routes
- **API Handlers**: `handler_*.go` files implement RESTful endpoints under `/v1`
- **Authentication**: API key-based auth via `Authorization: ApiKey <key>` header
- **Database**: SQLite via Turso, with sqlc-generated queries in `internal/database/`

### API Endpoints
- `POST /v1/users` - User registration
- `GET /v1/users` - Get authenticated user
- `POST /v1/notes` - Create note
- `GET /v1/notes` - List user's notes

### Frontend
Single-page application in `static/index.html` with vanilla JavaScript for interacting with the API.

## CI/CD Workflow

The project uses GitHub Actions (`.github/workflows/ci.yml`) which currently:
1. Triggers on pull requests to main
2. Sets up Go 1.23.0
3. Verifies Go version

Note: This is a learning project - tests are being implemented on the `addtests` branch.