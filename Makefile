migrate-up:
	migrate -path ./migrations -database 'sqlite3://store.db' up

migrate-down:
	migrate -path ./migrations -database 'sqlite3://store.db' down

migrate-create:
	migrate create -ext sql -dir migrations -seq create_foo_table