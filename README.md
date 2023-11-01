# Structure of Scrap Project
    │ scrap-market
    │   ├── main_service (scrapmarketbe)
    │   ├── notification_service (scrapmarketv2_notification_service)
    │   ├── inventory_service (scrapmarketv2_inventory_service)
    │   ├── proto (Package will push on github and download to source, link below)     
    │   ├── docker-compose.yml
[Repo proto](https://github.com/tuanpham197/test_repo)
## Config client in GRPC (Use scrapmarketbe for client)
### 1. File Config
  ```
  │ scrapmarketbe
  │   ├── cmd
  │   ├── composer
  │   ├── ...
  │   ├── internal
  |   |      ├── category
  |   |             ├── infras
  |   |                   ├── mysql
  |   |                   ├── rpc (rpc_hdl.go)
  |   ├── pkg
          ├── common
              ├── config.go
  ```
#### Folder cmd
- Init config port and server address
  ```
  flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"app_inventory:3300", // app_inventory is service name of inventory_service in docker-compose.yml
		"gRPC server address. Default: app_inventory:3300",
	)
  ```

### Folder composer
- Init client, service and controller
- File: grpc_client_composer.go
  - Init connect to GRPC server
- File: service_composer.go
  - Init interface contain method in api_handler.go
  - Init SQLRepo, RPCCLient, Service and Controller
### Folder pkg/common
- Define interface for config in folder cmd


## Config service in GRPC (Use scrapmarketv2_inventory_service for server RPC)
### 1. File Config
  ```
  │ scrapmarketbe
  │   ├── cmd
  │         ├── conf_comp.go
  |         ├── handler.go     
  │   ├── composer
  |         ├── service_composer.go
  │   ├── ...
  │   ├── internal
  |   |     ├── translate
  |   |             ├── controller
  |   |                   ├── api
  |   |                   ├── rpc (rpc_handler.go)
  |   ├── pkg
          ├── common
              ├── config.go
  ```
### Folder cmd
- File conf_comp.go
  - Init config port and server address
  ```
  func (c *config) InitFlags() {
    flag.IntVar(
      &c.grpcPort,
      "grpc-port",
      3300,
      "gRPC Port. Default: 3300",
    )

    flag.StringVar(
      &c.grpcServerAddress,
      "grpc-server-address",
      "localhost:3300",
      "gRPC server address. Default: localhost:3300",
    )
  }
  ```
- File handler.go
  - Start server GRPC with port config in conf_comp.go
### Folder internal/translate/controller/rpc (Client will call method in this file)
- Call service and return data

### Folder pkg/common
- Define interface for config in folder cmd

# 1 First Run: install package 
```
step1: go mod init sendo
step2: go mod tidy 
```
# 2 Start build image
 - Run command: make compose_build
# 3 Run project
 - Run command: make compose_up 

# 4 Generate docs swagger
 - swag init  OR swag init --parseDependency -g main.go 
 - If get error: command not found: swag, please run PATH=$(go env GOPATH)/bin:$PATH

# 5 Link docs API
 - http://localhost:4040/docs/index.html

 # Run with air
 Install air on macos
  1. curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
  2. Add alias air='~/.air' to your .bashrc or .zshrc
  3. check: air -v
 
 Install air in project
  go get -u github.com/cosmtrek/air

# Common cmd db
1. Dump database
```
docker exec containerId mysqldump -u root sendo_db > scrap.sql
```

2. Migration
Up
```
make migrate_up 
```
and
Down
```
make migrate_down
```
# Run unit test
After create test function, please add to Makefile
Example: 
```
go test -v -cover path_folder_service_test
```
Run test
```
make run_test
```


   

