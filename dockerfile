FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go" ]