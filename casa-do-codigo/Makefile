install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install-sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/casa_do_codigo?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/casa_do_codigo?sslmode=disable" down

generate:
	sqlc generate

dev-run:
	go run cmd/main.go

dev-infra-up:
	docker compose up -d

dev-infra-down:
	docker compose down