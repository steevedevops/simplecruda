build:
	@go build -o bin/simple_crud src/main.go

run: build
	@./bin/simple_crud

tidy:
	@go mod tidy

test:
	@go test -v ./...