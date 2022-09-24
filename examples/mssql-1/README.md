# Let's GO

https://github.com/denisenkom/go-mssqldb

https://blog.logrocket.com/using-sql-database-golang/

https://hevodata.com/learn/golang-mssql/

https://learn.microsoft.com/en-us/azure/azure-sql/database/connect-query-go?view=azuresql

go get github.com/denisenkom/go-mssqldb github.com/joho/godotenv/cmd/godotenv

# manual build

 go mod init main
 
```
go build
```

## docker
```
docker build --tag docker-go-wonbat .

docker run docker-go-wonbat
docker run --publish 8080:8080 docker-go-wonbat

docker compose up --build --remove-orphans
docker-compose -f docker-compose.test.yml up --build --remove-orphans

```
### REF:
- https://docs.docker.com/language/golang/build-images/
- https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
- https://github.com/olliefr/docker-gs-ping
- https://github.com/codefresh-contrib/golang-sample-app


