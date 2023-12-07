build.server:
	go build -o ./bin/main ./cmd/server/main.go

build.client:
	go build -o ./bin/main ./cmd/client/main.go

tidy:
	go mod tidy

proto:
	rm -f internal/pb/*.go
	protoc --proto_path=api/proto --go_out=internal/pb --go_opt=paths=source_relative \
	--go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative \
	api/proto/*.proto

docker.build:
	docker build -t go-grpc-http .

docker.up:
	docker compose up -d

docker.down:
	docker compose down

docker.clean:
	docker-compose kill && docker-compose rm -f
	docker rmi grpc_training_server:v1
	docker rmi grpc_training_client:v1

sqlc:
	docker run --rm -v /home/rubenadi/tuts/go-grpc-http:/src -w /src sqlc/sqlc generate

migrate.init:
	migrate create -ext sql -dir db/migrations -seq create_cars_table

migrate.up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/example_grpc?sslmode=disable" up

migrate.down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/example_grpc?sslmode=disable" down

evans:
	evans --host localhost --port 9090 -r repl

run.server:
	go run ./cmd/server/main.go

run.client:
	go run ./cmd/client/main.go

.PHONY: sqlc proto