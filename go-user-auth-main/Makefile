all: install

install:
	docker compose down
	docker image rm user-auth-api --force
	docker compose up -d --force-recreate
	@PATH=$(go env GOPATH)/bin:$$PATH GO111MODULE="on" && cd src && go mod download
	@echo "Proptuity service is running on http://localhost:3081"
	
test :
	go test ./... -v -cover

fmt:
	go fmt ./...

vet:
	go vet ./...

swag:
	swag init -parseDependency

generate:
	go generate ./...