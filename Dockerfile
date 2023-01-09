FROM golang:alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

WORKDIR /server
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

ENTRYPOINT /wait && ./main