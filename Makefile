createdb:
	go run lib/task/main/task.go db:create

migratenew:
	@read -p "Enter migration name: " userInput && \
    migrate create -ext sql -dir db/migration -seq "$$userInput"


migrateup:
	migrate -path db/migration -database "postgresql://Alone:Alone123!*@localhost:5432/exchange_production?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://Alone:Alone123!*@localhost:5432/exchange_production?sslmode=disable" -verbose down

checkfile:
	cat db/migration/*.sql

.PHONY: createdb migratenew migrateup migratedown checkfile