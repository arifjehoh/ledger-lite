.PHONY: run fmt test

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...