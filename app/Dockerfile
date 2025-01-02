FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init go-mysql-app
RUN go get -u github.com/go-sql-driver/mysql
RUN go build -o app

CMD ["./app"]
