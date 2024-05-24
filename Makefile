build:
	@go build -o bin/simple_crud src/main.go

run: build
	@./bin/simple_crud

test:
	@go test -v ./...