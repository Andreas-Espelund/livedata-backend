.PHONY: server
server:
	go run ./generated/cmd/livedata-backend-server/main.go --port=8080

.PHONY: generate
swagger:
	swagger generate server -A livedata-backend -f ./manifests/openApi/swagger.yaml -t ./generated && go mod tidy

build:
	go build -v ./generated/cmd/livedata-backend-server/main.go

docker-build:
	docker build -t livedata-backend-server .

docker-run:
	docker run -p 8080:8080 livedata-backend-server

