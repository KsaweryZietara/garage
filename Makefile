.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fix
fix:
	go mod tidy
	golangci-lint run --fix

.PHONY: web
web:
	cd web && npx expo start
