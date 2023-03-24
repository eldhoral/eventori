run: build
	@./builds/eldho/eventori

build:
	@go mod tidy
	@go build -o builds/eldho/eventori main.go

test:
	@go fmt ./...
	@go vet ./...
	@go test -v -coverprofile=coverage.out ./...

coverage:
	@go tool cover -html=coverage.out

migration:
	@go run . migrate

mysql:
	@docker-compose up