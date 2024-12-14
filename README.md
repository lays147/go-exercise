# Golang Challenge

To run the tests please use:
```shell
go test
```

To check code coverage please use:
```shell
go install gotest.tools/gotestsum@latest
gotestsum --format testname -- -coverprofile=coverage.txt -covermode=atomic ./... 
go tool cover -html=coverage.txt
```
