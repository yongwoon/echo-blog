# echo-blog

echo framework, mysql, gorm

## Set dev

- [Set develop env](./docs/dev.md)

## execute test code

```bash
ENVIRONMENT=test go test -v ./...
ENVIRONMENT=test go test ./...
```

## Swagger

- [local swagger](http://localhost:1323/swagger/index.html)

### How to use?

```bash
# Access to api container
docker-compose exec api bash

# run server
air -c .air.conf
```

then access to [swagger](http://localhost:1323/swagger/index.html)

### How to update swagger ?

1. update docs in controllers folder
2. in root path, execute commend as follow

```bash
swag i
```

## Reference

- [Installed packages](docs/packages.md)
- [sql-migration command](docs/sql_migrate.md)
- [Directory structure](docs/structure.md)
- [How to write swagger](https://github.com/swaggo/swag#api-operation)

## Links

- [project layout](https://github.com/golang-standards/project-layout)
- [golang echo realworld example app](https://github.com/xesina/golang-echo-realworld-example-app)
- [build web application with golang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/)
- [goconvey](http://goconvey.co/)
- [goconvey github](https://github.com/smartystreets/goconvey)
- [gconvey in docker?](https://github.com/smartystreets/goconvey/issues/449)
- [net/http status codes](http://golang.jp/pkg/http)
- [echo middleware test](https://github.com/labstack/echo/issues/659)
