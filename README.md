![ReadMe Banner](https://raw.githubusercontent.com/MathisVerstrepen/github-visual-assets/refs/heads/main/banner/PortfolioV2.png)

## Prerequisites

- Go 1.25 or later
- templ CLI tool
- Node.js (for Tailwind CSS)
- Air (for live reload)
- Docker with Docker Compose (for production deployment)

## Getting Started

1. Clone this repository

2. Copy the `.env.example` file to `.env`:

   ```bash
   cp .env.example .env
   ```

3. Install Go dependencies:

   ```bash
   go mod tidy
   ```

4. Install Air for live reload:

   ```bash
   go install github.com/air-verse/air@latest
   ```

5. Install templ CLI tool:

   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

6. Install Tailwind CSS dependencies:

   ```bash
   npm install
   ```

7. Run the development server:

   ```bash
   make dev
   ```

## Make Commands

| Command         | Description                                   |
| --------------- | --------------------------------------------- |
| `make dev`      | Start dev server with live reload (Air)       |
| `make css`      | Build Tailwind CSS                            |
| `make generate` | Build CSS + generate templ components         |
| `make build`    | Generate + compile binary to `./bin/portfolio`|
| `make run`      | Generate + run the server                     |
| `make tidy`     | Run `go mod tidy`                             |
| `make test`     | Run tests                                     |
| `make verify`   | Full check: fmt + tidy + test + vet + build   |

## Project Structure

```text
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
├── .dockerignore
├── portfolio.Dockerfile
└── docker-compose.yml
```

## Development

- Use `make dev` to start the development server with live reload.
- Add new routes in `main.go`.
- Create handlers in the `handlers/` directory.
- Add templ components in the `components/` directory.

## Local Production Build

Build and run the binary directly:

```bash
make build
ENV=prod PORT=8080 ROOT_DIR=$(pwd) \
./bin/portfolio
```

## Production Deployment With Docker Compose

Production deployment is a local Docker Compose build. There is no ApolloLaunch deploy configuration and no registry image is required.

1. On the target host, clone or update the repository.

2. Optionally set the host port in `.env`:

   ```bash
   HOST_PORT=8090
   ```

   The container always runs with `ENV=prod`, `PORT=8080`, and `ROOT_DIR=/app` from `docker-compose.yml`.

3. Build the production image locally:

   ```bash
   docker compose build --pull
   ```

4. Start or update the container:

   ```bash
   docker compose up -d
   ```

5. Check the deployment:

   ```bash
   docker compose ps
   curl http://localhost:${HOST_PORT:-8090}/ping
   ```

Useful operations:

```bash
docker compose logs -f
docker compose restart
docker compose down
```

To redeploy after pulling new code, rebuild locally and recreate the container:

```bash
docker compose build --pull
docker compose up -d
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
