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
 

