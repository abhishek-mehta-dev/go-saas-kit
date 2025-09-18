.PHONY: dev build run test tidy clean atlas-diff atlas-apply-dry atlas-apply atlas-lint

# === Config ===
MIGRATIONS_DIR := file://migrations

# Use DATABASE_URL from environment
# Make sure to export DATABASE_URL in your shell or load from .env
# Example:
# export DATABASE_URL="postgresql://user:password@host:port/dbname?sslmode=require&search_path=public"

# === Go commands ===

# Start server in dev mode with Air (hot reload)
dev:
	air

# Build production binary
build:
	go build -o bin/go-saas-kit ./cmd/app

# Run production binary
run:
	./bin/go-saas-kit

# Run all tests
test:
	go test ./... -v

# Clean & tidy dependencies
tidy:
	go mod tidy
	go mod vendor

# Remove build artifacts
clean:
	rm -rf bin tmp

# === Atlas migrations ===

# Generate a new migration from GORM models (timestamped by default)
atlas-diff:
	@read -p "Migration name (leave empty for timestamped name): " name; \
	if [ -z "$$name" ]; then \
		name=$$(date +%Y%m%d%H%M%S); \
	fi; \
	atlas migrate diff $$name --env gorm --dev-url "docker://postgres/15/dev"

# Apply migrations to DB (dry run)
atlas-apply-dry:
	atlas migrate apply --dir $(MIGRATIONS_DIR) --url $$DATABASE_URL --dry-run --revisions-schema "public"

# Apply migrations to DB
atlas-apply:
	atlas migrate apply --dir $(MIGRATIONS_DIR) --url $$DATABASE_URL --revisions-schema "public"

# Check migration safety / lint
atlas-lint:
	atlas migrate lint --env gorm --dev-url "docker://postgres/15/dev" --latest 1
