compose_up_rebuild:
	docker compose up --build --force-recreate
compose_up:
	docker compose up -d
compose_build:
	docker compose build --no-cache
migrate:
	migrate create -ext sql -tz Asia/Ho_Chi_Minh -dir db/migrations create_table_name_table
migrate_up:
	migrate -database "mysql://root:@tcp(localhost:3301)/sendo_db" -path db/migrations up

migrate_down:
	migrate -database "mysql://root:@tcp(localhost:3301)/sendo_db" -path db/migrations down
build_proto:
	protoc --proto_path=path_file_proto --go_out=folder_out --go_opt=paths=source_relative \
    --go-grpc_out=folder_out --go-grpc_opt=paths=source_relative \
    file_proto_name
run_test:
	go test -v -cover ./internal/translate/service
	go test -v -cover ./internal/category/service

# the name of the binary when built
BINARY_NAME=my-app

# remove any binaries that are built
clean:
	rm -f ./bin/$(BINARY_NAME)*

build-debug: clean
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o bin/$(BINARY_NAME)-debug main.go