build:
	@go build -o bin/api

run: build
	@./bin/api

seed:
	@go run scripts/seeds.go

test:
	@go test -v ./... -count=1
