# Set dev

### set env file

```bash
cp .env.sample .env
```

### build docker image

```bash
docker-compose build --no-cache
```

### run docker

```bash
docker-compose up
```

### datbase setting

```bash
# access to db container
docker-compose exec db bash

# acess mysql (just enter when asked 「Enter password」)
mysql -uroot -p
```

```sql
-- create db
CREATE DATABASE echo_blog_development CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE DATABASE echo_blog_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- exit mysql
exit
```

```bash
# exit container
exit
```

### migration

- access to api container

```bash
docker-compose exec api bash
```

- migration

```bash
sql-migrate up -config=config/database.yml -env="development"
```

- reference
  - [sql_migrate](./../api/docs/sql_migrate.md)

### run server

- access to api container

```bash
docker-compose exec api bash
```

- run

```bash
air -c .air.conf
```
