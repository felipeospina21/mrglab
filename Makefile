start:
	@echo "Starting app..."
	go run .

build:
	go build .

test:
	go test ./...

cover: 
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"
