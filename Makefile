build:
	go build ./cmd/nqg

clean:
	go clean

lint:
	golangci-lint run ./...

run:
	go run ./cmd/nqg

target:
	mkdir -p target

test: target
	go test -cover -coverprofile=./target/cover.out -v ./...

vulncheck:
	govulncheck ./...

.PHONY: build clean lint run test vulncheck
