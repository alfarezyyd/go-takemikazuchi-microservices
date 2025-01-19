# Gunakan shell bash
SHELL := /bin/bash

# Variabel
GO := go
RUN_CMD := $(GO) run cmd/server/main.go

# Command utama
.PHONY: run build clean test fmt vet lint wire gen inject

# Menjalankan aplikasi
default: run

run:
	$(RUN_CMD)

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

# Build injection
inject:
	wire gen ./cmd/injection/injector.go