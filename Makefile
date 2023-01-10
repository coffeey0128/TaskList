run:
	go run main.go
 
build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

doc:
	swag init

migrate-up:
	docker exec mysql mysql -uroot -psecret -e \
	"CREATE DATABASE IF NOT EXISTS gogolook;"
	sql-migrate up -config=./migration/dbconfig.yml --env=localhost

migrate-down:
	sql-migrate down -config=./migration/dbconfig.yml --env=localhost


seed-flush:
	# flush mysql
	docker exec mysql mysql -uroot -psecret -e \
	"SELECT CONCAT('TRUNCATE TABLE ', table_schema, '.', TABLE_NAME, ';') FROM INFORMATION_SCHEMA.TABLES \
	WHERE table_schema IN ('gogolook') AND TABLE_NAME != 'migrations'" | grep "gogolook*" | xargs -I {} docker exec mysql mysql -uroot -psecret -e {}
	# exec seeder
	go run ./cmd/seeder/main.go	

test:
	go test -v ./...

testf: seed-flush test

start: migrate-up seed-flush doc run

gen-mock:
	mockgen -source=./internal/task/repository.go -destination=./mock/task/repository.go -package=mock_task

gen-crud:
	go run ./code_gen/main.go && wire ./api . && make gen-mock
	
