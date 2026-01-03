# OrderAPI

OrderAPI — backend-сервис на Go для управления заказами.  
Проект реализован с разделением на слои (repository / service / HTTP), поддерживает REST API, PostgreSQL и корректное завершение работы (graceful shutdown).

## 🧩 Архитектура

Проект построен по принципам чистой архитектуры:

HTTP (handlers, router)

↓

Service (бизнес-логика)

↓

Repository (PostgreSQL)


Зависимости собираются вручную в `main.go` 

## 📦 Стек технологий

- Go 1.22+
- net/http
- PostgreSQL
- database/sql
- Docker (в процессе)
- REST API
- Context API
- Graceful shutdown

## 📁 Структура проекта


## 🗄️ Модель данных

### Таблица `orders`

```sql
CREATE TABLE orders (
    id          BIGSERIAL PRIMARY KEY,
    status      VARCHAR(32) NOT NULL,
    amount      BIGINT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
```
Статусы заказа

new

processing

completed

failed

1. Установить переменные окружения
   ```bash
   export POSTGRES_DSN="postgres://user:password@localhost:5432/order_db?sslmode=disable"
   export HTTP_PORT=8080
2. Запустить сервер
    ```bash
    go run cmd/api/main.go
