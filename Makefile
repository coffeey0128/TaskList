.PHONY : run build
run:
	go run main.go
 
build:
	GOOS=linux GOARCH=amd64 go build -o main main.go

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
	WHERE table_schema IN ('task_list') AND TABLE_NAME != 'migrations'" | grep "task_list*" | xargs -I {} docker exec mysql mysql -uroot -psecret -e {}
	# exec seeder
	go run ./cmd/seeder/main.go	

gen-mock:

gen-crud:
	go run ./code_gen/main.go
	