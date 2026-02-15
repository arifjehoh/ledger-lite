.PHONY: run fmt

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: fmt
fmt:
	go fmt ./...