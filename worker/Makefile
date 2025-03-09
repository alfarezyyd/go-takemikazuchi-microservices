# Gunakan shell bash
SHELL := /bin/bash

# Variabel
GO := go
MIGRATE_COMMAND := migrate
RUN_CMD := $(GO) run cmd/server/main.go
TABLE = $(TABLE)
NAME = $(NAME)
URL=mysql://root@tcp(127.0.0.1:3306)/go_takemikazuchi_microservices_workers
VERSION=$(VERSION)
PROTO_NAME=$(proto)

# Command utama
.PHONY: run build clean test fmt vet lint wire gen inject migrate migration-up gen


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

migration-up:
	$(MIGRATE_COMMAND) -database "$(URL)" -path ./migrations up

migration-down:
	$(MIGRATE_COMMAND) -database "$(URL)" -path ./migrations down

migration-create:
	$(MIGRATE_COMMAND) create -ext sql -dir ./migrations $(NAME)

migration-force:
	$(MIGRATE_COMMAND) -database "$(URL)" -path ./migrations force $(VERSION)
