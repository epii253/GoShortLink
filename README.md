# GoScanner - сервис сокращения ссылок

Минималистичный сервис на Go для создания и резолва коротких ссылок.
Построен на чистой слоистой архитектуре: handlers, services, repositories.

## Стек

- **Go** + **Gin** — HTTP сервер
- **GORM** + **PostgreSQL** — хранение данных, автомиграции при старте
- **Redis** — кеширование запросов на чтение
- **google/wire** — компайл-тайм dependency injection
- **Docker / docker-compose** — контейнеризация и оркестрация

## Структура проекта

```
.
├── cmd/app/
│   └── main.go                              # Точка входа
├── internal/
│   ├── application/
│   │   ├── contracts/
│   │   │   └── links_models.go              # DTO (запросы и ответы)
│   │   ├── repositories/
│   │   │   └── links_repo.go                # Интерфейс репозитория
│   │   └── services/
│   │       ├── links/
│   │       │   └── link_service.go          # Бизнес-логика
│   │       └── utilities/
│   │           └── url_convertor.go         # Генерация короткого кода
│   ├── controllers/
│   │   └── link_handlers/
│   │       └── handler.go                   # HTTP обработчики
│   ├── di/
│   │   ├── wire.go                          # Описание зависимостей (wire)
│   │   └── wire_gen.go                      # Сгенерированный DI код
│   ├── domain/
│   │   └── link_model.go                    # Доменная модель Link
│   ├── infrastructure/
│   │   ├── db_connection.go                 # Подключение к PostgreSQL, AutoMigrate
│   │   ├── db_cached.go                     # Кеширующий декоратор репозитория (Redis)
│   │   ├── db_redis_connection.go           # Подключение к Redis
│   │   ├── links_repo_db.go                 # Реализация репозитория через GORM
│   │   └── links_repo_memory.go             # In-memory реализация (для разработки)
│   └── settings/
│       └── config.go                        # Загрузка конфигурации из .env
├── Dockerfile
├── docker-compose.yml
└── .env.example
```

## Быстрый старт

### Через Docker (рекомендуется)

Скопируй файл с переменными окружения и заполни значения:

```bash
cp .env.example .env
```

Запусти сервис вместе с базой данных и Redis:

```bash
docker compose up --build
```

При старте приложение автоматически создаст таблицы в базе данных.

### Локально

Для локального запуска нужны запущенные PostgreSQL и Redis. Заполни `.env` и запусти:

```bash
go run ./cmd/app
```

## Конфигурация

Все параметры задаются через `.env` файл. Пример значений смотри в `.env.example`.

| Переменная          | Описание                             | По умолчанию  |
|---------------------|--------------------------------------|---------------|
| HOST                | Адрес сервера                        | 0.0.0.0       |
| PORT                | Порт сервера                         | 8080          |
| DB_HOST             | Хост базы данных                     | localhost     |
| DB_PORT             | Порт базы данных                     | 5432          |
| POSTGRES_USER       | Пользователь БД                      |               |
| POSTGRES_PASSWORD   | Пароль БД                            |               |
| POSTGRES_DB         | Имя базы данных                      |               |
| REDIS_ADDR          | Адрес Redis (host:port)              | localhost:6379|
| REDIS_PASSWORD      | Пароль Redis (пусто если нет)        |               |
| CACHE_TTL           | Время жизни кеша (например 2m, 30s)  | 2m            |

## Кеширование

Запросы на получение ссылки (`GetByLink`) кешируются в Redis. При записи и удалении
кеш инвалидируется автоматически. Если Redis недоступен при старте приложение не запустится.

## HTTP API

### POST /link

Создать короткую ссылку.

Тело запроса:

```json
{
  "full_url": "https://example.com/very/long/path"
}
```

Пример вызова:

```bash
curl -X POST http://localhost:8080/link \
  -H "Content-Type: application/json" \
  -d '{"full_url":"https://example.com/very/long/path"}'
```

Ответ (201 Created):

```json
{
  "url": "Ab3XyZ7"
}
```

### GET /link/:shortUrl

Редирект на оригинальный URL по короткому коду.

```bash
curl -v http://localhost:8080/link/Ab3XyZ7
```

Возвращает `302 Found` с заголовком `Location`.

## Разработка

После изменений в `internal/di/wire.go` нужно перегенерировать DI код:

```bash
wire ./internal/di
```

## Планы

- Тесты (unit и integration)
- Статистика переходов
- TTL для ссылок
- Rate limiting
