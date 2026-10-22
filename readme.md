# Go Web Application with PostgreSQL, Docker & golang-migrate

Простой бэкенд на Go с регистрацией пользователей, хэшированием паролей и автоматическими миграциями базы данных через golang-migrate. Проект полностью контейнеризирован.

---

## Технологический стек

- Язык: Go (net/http, database/sql, text/template)
- База данных: PostgreSQL 15
- Миграции: github.com/golang-migrate/migrate/v4
- Безопасность: golang.org/x/crypto/bcrypt
- Инфраструктура: Docker, Docker Compose

---

## Особенности реализации

- Автоматические миграции: приложение само применяет SQL-скрипты из папки migrations/ при старте.
- Безопасность: пароли хэшируются с помощью bcrypt.
- Изоляция: приложение и база запущены в отдельных Docker-контейнерах.
- Персистентность: данные базы сохраняются в Docker Volume.

---

## Структура проекта

.
├── cmd/app/main.go                  # Точка входа в приложение
├── internal/
│   ├── db/db.go                     # Подключение к БД и запуск миграций
│   └── handlers/auth.go             # Логика регистрации и хэширования
├── migrations/
│   ├── 000001_create_user_table.up.sql   # Накатывание миграции
│   └── 000001_create_user_table.down.sql # Откат миграции
├── web/templates/register.html      # Шаблон страницы регистрации
├── Dockerfile                       # Многоэтапная сборка образа
└── docker-compose.yml               # Конфигурация контейнеров
---

## Запуск приложения

1. Соберите и запустите проект:
   docker compose up --build

2. Откройте страницу регистрации в браузере:
   http://localhost:8080/register

---

## Проверка данных в базе

Подключитесь к базе данных через контейнер:
docker compose exec db psql -U root -d mydb

После этого можно посмотреть пользователей:
SELECT * FROM users;

---

## Остановка

- Остановить контейнеры (с сохранением данных):
  docker compose down

- Остановить с полной очисткой базы данных:
  docker compose down -v