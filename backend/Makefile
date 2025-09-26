deps:
	@echo "Installing dependencies..."
	go install github.com/daixiang0/gci@latest
	go mod tidy
	go mod download
.PHONY: deps

mock:
	go generate ./...
.PHONY: mock

build:
	docker-compose up --build -d
.PHONY: build

start:
	docker-compose up -d
.PHONY: start

stop:
	docker-compose down -v
.PHONY: stop

test:
	go test -count=1 -p 1 ./...
.PHONY: test

lint:
	go mod tidy
	go vet ./...
	gci write -s standard -s default -s "prefix(github.com/Lionel-Wilson/arritech-takehome/backend)" .
	gofumpt -l -w .
	wsl -fix ./... 2> /dev/null || true
	golangci-lint run $(p)
	go fmt ./...
.PHONY: lint