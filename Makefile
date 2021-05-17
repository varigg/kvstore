BINARY_NAME=kvstore
 
all: build test
 
build:
	go build -o ${BINARY_NAME} main.go
 
test:
	go test -v -coverprofile=coverage.out ./...
 
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}