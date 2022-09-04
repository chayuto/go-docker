# Let's GO

# manual build

```
go build
```

## docker
```
docker build --tag docker-go-wonbat .

docker run docker-go-wonbat
docker run --publish 8080:8080 docker-go-wonbat
```
### REF:
- https://docs.docker.com/language/golang/build-images/
- https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
- https://github.com/olliefr/docker-gs-ping

