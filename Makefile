# Gunakan shell bash
SHELL := /bin/bash

# Variabel
GO := go

PROTO_NAME=$(proto)

# Command utama
.PHONY: build clean test fmt vet lint gen

# Build aplikasi
build:
	$(GO) build -o bin/app cmd/server/main.go

# Membersihkan file hasil build
clean:
	rm -rf bin/

# Menjalankan unit test
test:
	$(GO) test ./... -cover

# Format kode
dfmt:
	$(GO) fmt ./...

# Static analysis
vet:
	$(GO) vet ./...

# Linting menggunakan golangci-lint
lint:
	golangci-lint run ./...

# Menjalankan semua check sekaligus
check: fmt vet lint test

gen:
	@protoc \
	--proto_path=protobuf "protobuf/$(PROTO_NAME).proto" \
	--go_out=common/genproto/$(PROTO_NAME) --go_opt=paths=source_relative \
  	--go-grpc_out=common/genproto/$(PROTO_NAME) --go-grpc_opt=paths=source_relative