.PHONY : run 
run:
	go run main.go
	
.PHONY : build 
build:
	GOOS=linux GOARCH=amd64 go build -o main main.go