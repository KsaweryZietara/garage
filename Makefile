.PHONY: run
run:
	docker compose up -d

.PHONY: stop
stop:
	docker compose down

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
