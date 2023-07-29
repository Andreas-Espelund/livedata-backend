.PHONY: server
server:
	go run ./cmd/myproject-server/main.go --port=8080

.PHONY: generate
generate:
	swagger generate server -A livedata-backend -f ./swagger.yaml