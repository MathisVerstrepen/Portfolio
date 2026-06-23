FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git nodejs npm

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

RUN npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify
RUN templ generate -path ./components

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /portfolio .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN addgroup -S portfolio && adduser -S portfolio -G portfolio

WORKDIR /app
ENV ENV=prod \
    PORT=8080 \
    ROOT_DIR=/app

COPY --from=builder --chown=portfolio:portfolio /portfolio ./portfolio
COPY --from=builder --chown=portfolio:portfolio /app/assets ./assets

USER portfolio

EXPOSE 8080

CMD ["./portfolio"]
