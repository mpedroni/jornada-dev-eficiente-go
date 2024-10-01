migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/casa_do_codigo?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/casa_do_codigo?sslmode=disable" down