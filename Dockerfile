FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /myproject ./cmd/myproject-server/main.go

EXPOSE 8080

CMD [ "/myproject", "--port=8080" ]