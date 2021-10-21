.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	gofmt -w ./
	golangci-lint run --config ./build/ci/.golangci.yml ./...

.PHONY: test
test:
	make generate
	go test ./... --count=1

.PHONY: uo_local
up_local:
	docker-compose -f build/docker-compose.yaml up -d

.PHONY: build
build:
	go build cmd/main.go

