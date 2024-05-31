build:
	@go build -o bin/simple_crud src/main.go

run: build
	@./bin/simple_crud

tidy:
	@go mod tidy

test:
	@go test -v ./...

# db operations
migrations:
	@go run src/main.go db migrate

migrationsback:
	@go run src/main.go db rollback

migrationsstatus:
	@go run src/main.go db status

# make createmigrationgo name_file=init
createmigration_go:
	@go run src/main.go db create_go $(name)

createmigration_sql:
	@go run src/main.go db create_sql $(name)

migrationshelp:
	@go run src/main.go db