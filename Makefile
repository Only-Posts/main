build:
	go build -o bin/main ./cmd

test:
	go test ./...

migrate:
	cd .\migrations\ && \
	goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up

migrate-status:
	cd .\migrations\ && \
    goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" status