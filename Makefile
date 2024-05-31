build:
	@go build -o bin/simple_crud src/main.go

run: build
	@./bin/simple_crud

tidy:
	@go mod tidy

test:
	@go test -v ./...

# db operations
initmigrations: # cria as tabelas de migracao iniciais 
	@go run src/main.go db init

migrations:
	@go run src/main.go db migrate

rollbackmigrations:
	@go run src/main.go db rollback

statusmigrations:
	@go run src/main.go db status

# make createmigrationgo name_file=init
creategomigration:
	@go run src/main.go db create_go $(name)

createsqlmigration:
	@go run src/main.go db create_sql $(name)

helpmigrations:
	@go run src/main.go db