# GoShortLink — микросервис для сокращённых ссылок

ShortLink — минималистичный микросервис на Go для создания и резолва сокращённых ссылок.
Проект находится на ранней стадии, но уже содержит архитектурный каркас (handlers → services → repositories)
и готов к расширению.

---

## Структура проекта

```
.
├── cmd/
│   └── app/
│       └── main.go                 # Точка входа, wiring зависимостей
├── go.mod
└── internal/
    ├── application/
    │   ├── repositories/
    │   │   └── links_repo.go       # Интерфейс репозитория
    │   └── services/
    │       └── links/
    │           ├── link_service.go # Бизнес-логика
    │           └── links_models.go # Модели / DTO
    ├── controllers/
    │   └── link_handlers/
    │       └── handler.go          # HTTP handlers (Gin)
    ├── infrastructure/
    │   └── links_repo_impl.go      # In-memory реализация репозитория
    └── settings/
        └── config.go               # Загрузка конфигурации (.env, fallback)
```

---

## Возможности

- Создание короткой ссылки по длинному URL
- Редирект по короткому коду
- Чистое разделение ответственности (handlers / services / repositories)
- Dependency Injection через конструкторы
- Потокобезопасная работа (stateless services)

---

## Быстрый старт

### 1. Запуск

```bash
go run ./cmd/app
```

По умолчанию сервер запускается на порту **8080**.

---

## HTTP API

### POST `/link`

Создание короткой ссылки.

**Тело запроса:**

```json
{
  "full_url": "https://example.com/very/long/path"
}
```

**Пример:**

```bash
curl -X POST http://localhost:8080/link   -H "Content-Type: application/json"   -d '{"full_url":"https://example.com/very/long/path"}'
```

**Ответ (пример):**

```json
{
  "short_url": "Ab3XyZ",
  "full_url": "https://example.com/very/long/path"
}
```

---

### GET `/link/:shortUrl`

Редирект по короткому коду.

```bash
curl -v http://localhost:8080/link/Ab3XyZ
```

Ожидаемый результат:
- `302 Found`
- заголовок `Location` с оригинальным URL

---

## Конфигурация

Конфигурация загружается при старте приложения:
- `.env`
- fallback: `.env.example`

Загрузка выполняется **один раз** (`sync.Once`).

---

## Текущие ограничения

- Используется **in-memory репозиторий** (данные теряются при перезапуске)
- Отсутствуют тесты
- Нет сбора статистики и аналитики

---

## Точки роста

- Добавить unit и integration тесты
- Добавить сбор статистики переходов
- Добавить поддержку PostgreSQL (вместо in-memory)
- Кеширование (Redis)
- Rate limiting и базовая защита от abuse

---

## Идеи для развития

- TTL для ссылок
- Статистика по IP / User-Agent
- Асинхронная обработка событий (channels / worker pool)
- Docker / docker-compose
- Health-check endpoints
