.PHONY: dev css generate build run tidy test verify

dev:
	air

css:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

generate: css
	templ generate -path ./components

build: generate
	go build -o ./bin/portfolio .

run: generate
	go run .

tidy:
	go mod tidy

test:
	go test ./...

verify: generate
	gofmt -w ./components ./handlers .
	go mod tidy
	go test ./...
	go vet ./...
	go build ./...
