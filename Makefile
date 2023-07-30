.PHONY: server
server:
	go run ./generated/cmd/livedata-backend-server/main.go --port=8080

.PHONY: generate
swagger:
	swagger generate server -A livedata-backend -f ./manifests/openApi/swagger.yaml -t ./generated && go mod tidy
