![ReadMe Banner](https://raw.githubusercontent.com/MathisVerstrepen/github-visual-assets/refs/heads/main/banner/PortfolioV2.png)

## Prerequisites

- Go 1.25 or later
- templ CLI tool
- Node.js (for Tailwind CSS)
- Air (for live reload)
- Docker (optional, for containerization)

## Getting Started

1. Clone this repository

2. Copy the `.env.example` file to `.env`:
   ```
   cp .env.example .env
   ```

3. Install Go dependencies:
   ```
   go mod tidy
   ```

4. Install Air for live reload:
   ```
   go install github.com/air-verse/air@latest
   ```

5. Install templ CLI tool:
   ```
   go install github.com/a-h/templ/cmd/templ@latest
   ```

6. Install Tailwind CSS dependencies:
   ```
   npm install
   ```

7. Run the development server:
   ```
   make dev
   ```

## Make Commands

| Command         | Description                                   |
| --------------- | --------------------------------------------- |
| `make dev`      | Start dev server with live reload (Air)       |
| `make css`      | Build Tailwind CSS                            |
| `make generate` | Build CSS + generate templ components         |
| `make build`    | Generate + compile binary to `./bin/portfolio` |
| `make run`      | Generate + run the server                     |
| `make tidy`     | Run `go mod tidy`                             |
| `make test`     | Run tests                                     |
| `make verify`   | Full check: fmt + tidy + test + vet + build   |

## Project Structure

```
.
├── assets/
│   ├── css/
│   ├── favicon/
│   ├── fonts/
│   └── img/
├── components/       # templ components
├── handlers/         # route handlers
├── main.go           # entry point
├── .air.toml         # Air live reload config
├── tailwind.config.js
├── package.json
├── portfolio.Dockerfile
└── docker-compose.yml
```

## Development

- Use `make dev` to start the development server with live reload.
- Add new routes in `main.go`.
- Create handlers in the `handlers/` directory.
- Add templ components in the `components/` directory.

## Building for Production

Build and run:
```
make build
./bin/portfolio
```

## Docker

Build and run with Docker:
```
docker compose build
docker compose up
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
