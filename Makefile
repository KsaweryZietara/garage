.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test: ## run Go tests
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fix
fix:
	go mod tidy
	golangci-lint run --fix
