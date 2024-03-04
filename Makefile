migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5460 sslmode=disable" goose up migrations/psql
migrate-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5460 sslmode=disable" goose down migrations/psql
run:
	@go run cmd/app/main.go
run-docker:
	@docker-compose up -d
	@sleep 2
	@GOOSE_DRIVER=clickhouse GOOSE_DBSTRING="tcp://127.0.0.1:9000?username=default&password=default" goose -dir migrations/click up 