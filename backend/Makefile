run:
	go run ./cmd/main.go

build:
	go build -o ./bin/bebrah ./cmd/main.go

generate-swag:
	swag init --parseDependency -g route.go -d ./app/server/