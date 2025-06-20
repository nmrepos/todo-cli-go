# Todo App in Go

This project implements a simple todo list with a web interface and REST API built with the Gin framework. A GitHub Actions workflow builds, tests and deploys a Docker image for dev and prod environments.

## Project Setup

Requires Go 1.20 or higher.

```bash
# build and run web server
go build ./cmd/server
./server
```

Open <http://localhost:8080> to use the browser interface.

## CI Pipeline

The workflow at `.github/workflows/ci.yml` performs:

1. **Build** – Compile the server in `cmd/server`.
2. **Test** – Run unit tests in the `todo` package.
3. **Lint** – Ensure code is formatted with `gofmt`.
4. **Docker** – Build and push an image to GitHub Container Registry.
5. **Deploy** – Separate jobs for `dev` (on pushes to `develop`) and `prod` (manual).

## Docker

```bash
docker build -t todo-app .
docker run -p 8080:8080 todo-app
```

Then visit `http://localhost:8080`.

## Jenkins

Not used; all CI/CD is handled via GitHub Actions.
