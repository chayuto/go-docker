# TCP server

https://github.com/mactsouk/opensource.com/blob/master/concTCP.go

# manual build

```
go mod init main
go mod tidy

go run .\main.go 5000

Test-NetConnection -ComputerName localhost -Port 5000
```

## docker
```
docker build --tag docker-go-wonbat .

docker run docker-go-wonbat
docker run --publish 8080:8080 docker-go-wonbat

docker compose up --build --remove-orphans
docker-compose -f docker-compose.test.yml up --build --remove-orphans

```
