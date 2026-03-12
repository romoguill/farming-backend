dev:
	air

start:
	go run cmd/main.go

db:
	docker compose up -d

db-down:
	docker compose down

migrateup:
	migrate -path internal/database/migration -database "postgresql://root:password@localhost:5432/farming-db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/database/migration -database "postgresql://root:password@localhost:5432/farming-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: start dev sqlc db db-down migrateup migratedown