SHELL := /bin/bash
USER_DB=root
PASS_DB=root
HOST=app-database
DSN=host=$(HOST) port=5432 user=$(USER_DB) password=$(PASS_DB)  dbname=sistema-escolar

create-migration:
	goose -dir internal/infra/database/postgres/migrations create $(name) sql
migrate:
	goose -dir internal/infra/database/postgres/migrations -table _db_version postgres '$(DSN)' up
migrate-rollback:
	goose -dir internal/infra/database/postgres/migrations -table _db_version postgres '$(DSN)' down
create-seed:
	goose -dir internal/infra/database/postgres/seeds create $(name) sql
seed:
	goose -dir internal/infra/database/postgres/seeds -table _db_seeds postgres '$(DSN)' up
sql:
	sqlc generate		