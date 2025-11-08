# Lucy

**Lucy** is a modern file sharing and storage platform built with **Go** and **TypeScript**.

> This project is currently under development.

---

## Features

See the [TODO.md](TODO.md) file for details.

---

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL
- [Goose](https://github.com/pressly/goose) for database migrations
- (Optional) [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Taskfile](https://taskfile.dev/)

---

### Clone the Repository

```bash
git clone https://github.com/primekobie/lucy.git
cd lucy
```

### Environment Setup

Copy the example environment file and update values as needed:

```bash
cp .env.example .env
```

---

### Running with Docker

```bash
docker compose up --build
```

Once everything starts up, the API should be available at:
**[http://localhost:8080](http://localhost:8080)**

---

## Project Structure

```sh
lucy/
├── client # the frontend application built with solid.js
├── cmd
│   └── api # entry point for the backend service
├── http # http test files
├── internal
│   ├── handlers # Handlers and middleware
│   ├── mailer # Email and templats
│   ├── models # objects used in the application
│   ├── postgres # database logic
│   └── services # business logic
└── migrations # database schema migrations (Goose)
```

---

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---
