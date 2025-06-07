DB_URL=postgres://user:pass@localhost:5432/mydb?sslmode=disable

run: build
	./bin/indexer -pg-url "$(DB_URL)"

build:
	mkdir -p bin
	go build -o ./bin/indexer ./cmd/indexer

format:
	goimports -w .

db:
	docker compose up -d
