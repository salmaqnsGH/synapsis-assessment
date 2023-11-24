createmigrate:
	migrate create -ext sql -dir migration -seq init_schema

migrateup:
	migrate -path migration -database "postgresql://root:secret@localhost:5455/shopping?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:secret@localhost:5455/shopping?sslmode=disable" -verbose down

migrateuptest:
	migrate -path migration -database "postgresql://root:secret@localhost:5455/shopping_test?sslmode=disable" -verbose up

migratedowntest:
	migrate -path migration -database "postgresql://root:secret@localhost:5455/shopping_test?sslmode=disable" -verbose down

test:
	go clean -testcache
	go test ./test

run:
	go run main.go

.PHONY: createmigrate migrateup migratedown migrateuptest migratedowntest test run