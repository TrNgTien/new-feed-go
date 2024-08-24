.PHONY: build
build:
	go build -o cmd/new-feed-go/main.go

.PHONY: run
run:
	go run cmd/new-feed-go/main.go

.PHONY: swagger
swagger:
	swag init -g cmd/main.go \
	--exclude ./internal/ \
	-o ./docs/

.PHONY: dc-down
dc-up:
	docker-compose -f deployments/docker-compose.dev.yml up -d

.PHONY: dc-down
dc-down:
	docker-compose -f deployments/docker-compose.dev.yml down
