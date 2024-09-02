VERSION := $(shell cat VERSION)
PROJECT_NAME:=new-feed-go

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build \
		-ldflags "-X main.version=$(VERSION) -X 'app/build.User=$(id -u -n)' -X 'app/build.Time=$(date)'" \
		-o build/$(PROJECT_NAME)_linux_amd64 cmd/$(PROJECT_NAME)/*.go

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build \
		-ldflags "-X main.version=$(VERSION) -X 'app/build.User=$(id -u -n)' -X 'app/build.Time=$(date)'" \
		-o build/$(PROJECT_NAME)_linux_arm64 cmd/$(PROJECT_NAME)/*.go

.PHONY: generate
generate:
	wire internal/wiring/wire.go

.PHONY: build
build:
	go build -o cmd/${PROJECT_NAME}/*.go

.PHONY: run
run: lint generate
	go run cmd/${PROJECT_NAME}/*.go new-feed-go

.PHONY: dc-down
dc-up:
	docker-compose -f deployments/docker-compose.dev.yml up -d

.PHONY: dc-down
dc-down:
	docker-compose -f deployments/docker-compose.dev.yml down

.PHONY: clean
clean:
	rm -rf build/

.PHONY: lint
lint:
	golangci-lint run ./... 