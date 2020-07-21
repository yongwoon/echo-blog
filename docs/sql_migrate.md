# SQL mygrate

```bash
sql-migrate --help
sql-migrate new --help
```

## migration 生成

```bash
sql-migrate new -config=config/database.yml -env="development" [NAME]
```

## migration

```bash
sql-migrate up -config=config/database.yml -env="development"
```

## rollback

```bash
sql-migrate down -config=config/database.yml -env="development" -limit=1
```

## status

```bash
sql-migrate status -config=config/database.yml -env="development"
```
