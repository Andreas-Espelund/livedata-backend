FROM golang:1.20

# Set directory
WORKDIR /app

COPY go.* ./

RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags '-s' -o api-server generated/cmd/livedata-backend-server/main.go

EXPOSE 8080

CMD ["/app/api-server", "--port=8080"]
