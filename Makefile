create-migration:
	@migrate create -ext sql -dir db/migrations/ "${NAME}" && echo ">> new migration created"

migrate:
	@printf ">> Running migrations\n"
	@migrate -database "postgres://hellofreshdev:hellofreshdev@localhost:5432/todo?sslmode=disable" -path db/migrations up