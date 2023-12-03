build.server:
	go build -o ./bin/main ./cmd/server/main.go

build.client:
	go build -o ./bin/main ./cmd/client/main.go

proto:
	protoc --proto_path=api/proto --go_out=internal/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
		api/proto/*.proto

docker.build:
	docker build -t go-grpc-http .

docker.up:
	docker compose up -d

docker.down:
	docker compose down

sqlc:
	docker run --rm -v /home/rubenadi/tuts/go-grpc-http:/src -w /src sqlc/sqlc generate

migrate.init:
	migrate create -ext sql -dir db/migrations -seq create_cars_table

migrate.up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/example_grpc?sslmode=disable" up

migrate.down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/example_grpc?sslmode=disable" down

evans:
	$ docker run --rm -v "/home/rubenadi/tuts/go-grpc-http:/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
      --host localhost \
      --port 9090 \
      repl

run.dev:
	go run ./cmd/server/main.go

run.prod:
	export APP_ENV=production && go run ./cmd/server/main.go

.PHONY: build.server build.client proto docker.build docker.up docker.down sqlc migrate migrate.up migrate.down run.dev run.prod