GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init

# init tool
init:
	GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	go install github.com/nicksnyder/go-i18n/v2/goi18n@latest

.PHONY: api
# generate api
api:
	$(shell protoc --proto_path=./api \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:./api \
     	       --go-http_out=paths=source_relative:./api \
     	       --go-grpc_out=paths=source_relative:./api \
     	       --go-errors_out=paths=source_relative:./api \
     	       --validate_out=paths=source_relative,lang=go:./api \
    	       --openapi_out=fq_schema_naming=true,default_response=false:./api \
    	       $(API_PROTO_FILES))

.PHONY: build
# build
build:
	mkdir -p bin/
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	GOFLAGS=-mod=mod go generate ./...

.PHONY: wire
# generate DI code
wire:
	wire ./...

.PHONY: all
# generate all
all:
	make api;
	make generate;
	make i18n;

.PHONY: lint
# lint
lint:
	golangci-lint run

.PHONY: i18n
# generate i18n
i18n:
	cd i18n && goi18n merge active.*.toml translate.*.toml && go run internal/i18n/generate.go

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
