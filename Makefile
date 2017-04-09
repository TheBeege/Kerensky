SSHKEY=""
DEPLOY_HOST="localhost"

REPO="github.com/TheBeege/Kerensky"
BINARY_NAME="Kerensky"
BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`
GOLDFLAGS="-X main.branch=$(BRANCH) -X main.commit=$(COMMIT)"

#PACKAGES="models"

all: test build

setup:
	@echo "=== setup ==="
	@go get -u "github.com/golang/lint/golint"

vet:
	@echo "==== go vet ==="
	@go vet ./...

lint:
	@echo "==== go lint ==="
	@golint ./**/*.go

fmt:
	@echo "=== go fmt ==="
	@go fmt ./...

install: test
	@echo "=== go install ==="
	@go install -ldflags=$(GOLDFLAGS)

build: test
	@echo "=== go build ==="
	@mkdir -p bin/
	@go build -ldflags=$(GOLDFLAGS) -o bin/${BINARY_NAME}

test: fmt vet lint
	@echo "=== go test ==="
	@go test ./... -cover

deploy: test
	# Compile
	@mkdir -p bin/
	GOARCH=amd64 GOOS=linux go build -ldflags=$(GOLDFLAGS) -o bin/${BINARY_NAME}
	# Copy binaries
	@scp bin/${BINARY_NAME} $(DEPLOY_HOST):~/
	# Cleanup binaries
	@rm bin/${BINARY_NAME}

clean:
	@rm -rf bin/* pkg/*

.PHONY: setup vet lint fmt install build test deploy clean
