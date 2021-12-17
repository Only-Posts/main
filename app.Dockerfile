FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get github.com/pressly/goose/v3/cmd/goose

COPY . .

RUN go build -o /only-post-api cmd/main.go

CMD [ "/only-post-api", "run" ]