Для запуска БД из Docker изспользуем команду

```bash
docker run --name=todo-db -e POSTGRES_PASSWORD="qwerty" -p 5432:5432 -d --rm postgres
```

Для миграций

```bash
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

Для отката миграций

```bash
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
```

****