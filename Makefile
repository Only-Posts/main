build:
	go build -o bin/main ./cmd

test:
	go test ./...

migrate:
	goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up
