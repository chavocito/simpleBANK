#!make
include .env

# Define the Go compiler and its flags.
GO := go
GOFLAGS := -v

# Define the entc command for generating the schema.
ENTC := entc

# Define the name of your Go binary.
BINARY_NAME := myapp

# Define the name of your database migration directory.
MIGRATIONS_DIR := migrations

# Default target.
all: build

# Build the Go binary.
build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

# Run database migrations (up).
migrate-up:
	$(ENTC) migrate --path $(MIGRATIONS_DIR) --database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" up

# Reverse the last database migration (down).
migrate-down:
	$(ENTC) migrate --path $(MIGRATIONS_DIR) --database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable" down

# Create a new database migration.
create-migration:
	$(ENTC) create migration -t $(MIGRATIONS_DIR)

# Drop the database (use with caution).
drop-db:
	$(GO) run ./scripts/drop_db.go

# Run the Go application.
run:
	$(GO) run $(GOFLAGS) ./cmd/$(BINARY_NAME)

# Clean up the generated binary.
clean:
	rm -f $(BINARY_NAME)

.PHONY: all build migrate-up migrate-down create-migration drop-db run clean

