.DEFAULT_GOAL := build/web

# ========================================================================== #
# HELPERS
# ========================================================================== #

# help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFULE_LIST} | column -t -s ':' | sed -e 's/ˆ/ /'

.PHONY: confirm
confirm:
	@if [ -t 0 ]; then \
		echo 'Are you sure? [y/N]' && read ans && [ "$${ans:-N}" = "y" -o "$${ans:-N}" = "Y" ] || { echo "Aborted"; exit 1; }; \
	else \
		echo "Error: Cannot prompt for confirmation in non-interactive environment"; exit 1; \
	fi

# ========================================================================== #
# DEVELOPMENT
# ========================================================================== #

# run/web: run the cmd/web application
.PHONY: run/web
run/web:
	go run ./cmd/web/ -dsn="${SNIPPETBOX_DB}"

# db/psql: connect to the database usin psql
.PHONY: db/psql
db/psql:
	psql ${SNIPPETBOX_DB}

# db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations/ ${name}

# db/migrations/up: apply all database migrations 
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations/ -database "${SNIPPETBOX_DB}" up

# ========================================================================== #
# QUALITY CONTROL
# ========================================================================== #

# audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ========================================================================== #
# BUILD
# ========================================================================== #

current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

.PHONY: build/web
build/web:
	@echo 'Building cmd/web...'
	go build -ldflags=${linker_flags} -o=./bin/web ./cmd/web/

.PHONY: build/web/linux
build/web/linux:
	@echo 'Building cmd/web...'
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/web ./cmd/web/
