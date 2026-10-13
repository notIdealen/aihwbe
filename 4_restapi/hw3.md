# Упражнение 3. PostgreSQL и репозиторий

## Цель

Перейти от хранения в памяти к реальному SQL-хранилищу.

## Используемые технологии

- PostgreSQL;
- Docker;
- `pgx` или `database/sql`;
- миграции;
- SQL-запросы.

Рекомендую:

```bash
go get github.com/jackc/pgx/v5
```

Для миграций:

```bash
go get -u github.com/pressly/goose/v3
```

или:

```bash
go get -u github.com/golang-migrate/migrate/v4
```

## Запусти PostgreSQL в Docker

Например:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=app -p 5432:5432 -d postgres:16
```

Или сразу в `docker-compose.yml`:

```yaml
services:
  postgres:
    image: postgres:16
    environment:
      POSTGRES_USER: app
      POSTGRES_PASSWORD: app
      POSTGRES_DB: app
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

## SQL-схема

Создай таблицу `tasks`.

```sql
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'todo',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_created_at ON tasks(created_at);
```

## Структуры

### Конфигурация БД

```go
type DBConfig struct {
    DSN          string
    MaxOpenConns int32
    MaxIdleConns int32
    ConnTimeout  time.Duration
}
```

### Репозиторий

```go
type TaskRepository interface {
    Create(ctx context.Context, task Task) error
    GetByID(ctx context.Context, id string) (Task, error)
    List(ctx context.Context, filter TaskFilter) ([]Task, error)
    Update(ctx context.Context, task Task) error
    Delete(ctx context.Context, id string) error
}
```

### PostgreSQL-реализация

```go
type PostgresTaskRepository struct {
    pool *pgxpool.Pool
}
```

### Фильтр

```go
type TaskFilter struct {
    Status string
    Search string
    Limit  int
    Offset int
}
```

## Что нужно сделать

1. Убери `MemoryTaskStore`.
2. Реализуй `PostgresTaskRepository`.
3. Подключи пул соединений через `pgxpool`.
4. Добавь миграции.
5. Обнови `main.go`:
   - подключай БД;
   - передавай репозиторий в сервис;
   - закрывай соединения при остановке.

## SQL-запросы, которые нужно написать

### Вставка

```sql
INSERT INTO tasks (id, title, description, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);
```

### Выборка по ID

```sql
SELECT id, title, description, status, created_at, updated_at
FROM tasks
WHERE id = $1;
```

### Список

```sql
SELECT id, title, description, status, created_at, updated_at
FROM tasks
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
```

### Обновление

```sql
UPDATE tasks
SET title = $2,
    description = $3,
    status = $4,
    updated_at = now()
WHERE id = $1;
```

### Удаление

```sql
DELETE FROM tasks
WHERE id = $1;
```

## Критерии готовности

- Данные сохраняются после перезапуска сервиса.
- Есть миграции вверх и вниз.
- Используются параметризованные запросы.
- Нет SQL-инъекций.
- Сервер корректно подключается к БД.
- Если БД недоступна, сервер возвращает ошибку или не стартует.

---
