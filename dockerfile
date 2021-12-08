# Переписать нормально
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /only-post-api cmd/main.go

EXPOSE 8080

CMD [ "/only-post-api", "run" ]