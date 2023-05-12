POSTGRESQL_URL=postgres://root:secret@localhost:5432/todo # it's local url db, i'm not foolish =)

migrate-create:
	migrate create -ext sql -dir migrations/ -seq 'migrate_name'
.PHONY: migrate-create

migrate-up:
	migrate -database '$(POSTGRESQL_URL)?sslmode=disable' -path migrations/ up
.PHONY: migrate-up

migrate-down:
	migrate -database '$(POSTGRESQL_URL)?sslmode=disable' -path migrations/ down
.PHONY: migrate-down

swaggo:
	swag init -g **/**/*.go

run:
	go run ./cmd/app/main.go
