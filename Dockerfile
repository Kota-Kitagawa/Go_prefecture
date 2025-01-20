FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
