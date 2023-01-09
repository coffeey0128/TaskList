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

