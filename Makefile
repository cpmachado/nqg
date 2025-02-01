
build:
	go build

clean:
	go clean

lint:
	golangci-lint run .

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
