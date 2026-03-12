dev:
	air

start:
	go run cmd/main.go

migrateup:
	migrate -path internal/database/migration -database "postgresql://root:password@localhost:5432/farming-db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/database/migration -database "postgresql://root:password@localhost:5432/farming-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: start dev sqlc