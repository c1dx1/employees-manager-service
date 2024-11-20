# Employee Management System

## Описание

Система управления сотрудниками, включающая один микросервис и API Gateway. Основной функционал:

- Добавление сотрудника.
- Удаление сотрудника.
- Обновление данных сотрудника.
- Получение списка сотрудников компании.

Для взаимодействия между компонентами используется **gRPC**, а для внешнего интерфейса — **REST API** через **API Gateway**.

## Стек технологий

- **Go** — основной язык разработки.
- **gRPC** — взаимодействие между сервисами.
- **Gin** — реализация REST API Gateway.
- **PostgreSQL** — база данных.
- **Docker** и **docker-compose** — контейнеризация.
- **Postman** — тестирование API.
- **Миграции** — управление схемой базы данных (golang-migrations).

### Доступ:

- API Gateway доступен по адресу: [http://localhost:8080](http://localhost:8080).
- gRPC-сервис доступен на порту :50051.

---

## Примеры использования API

### 1. Добавление сотрудника

**Запрос**:
```json
POST /employees
Content-Type: application/json

{
  "name": "John",
  "surname": "Doe",
  "phone": "+123456789",
  "company_id": 1,
  "passport": {
    "type": "ID",
    "number": "12345678"
  },
  "department": {
    "name": "HR",
    "phone": "+987654321"
  }
}
```

**Ответ**:
```json
{
  "id": 1
}
```

---

### 2. Удаление сотрудника

**Запрос**:
```json
DELETE /employees
Content-Type: application/json

{
  *"id": 1
}
```

**Ответ**:
```json
{
  "success": "Success"
}
```

---

### 3. Получение списка сотрудников компании (есть возможность выборки по отделу компании)

**Запрос**:
```json
GET /employees
Content-Type: application/json

{
  *"company_id": 1,
  "department": {
    "name": "HR"
  }
}
```

**Ответ**:
```json
{
  "employees": [
    {
      "id": 1,
      "name": "John",
      "surname": "Doe",
      "phone": "+123456789",
      "company_id": 1,
      "passport": {
        "type": "ID",
        "number": "12345678"
      },
      "department": {
        "name": "HR",
        "phone": "+987654321"
      }
    }
  ]
}
```

---

### 4. Обновление данных сотрудника

**Запрос**:
```json
PUT /employees
Content-Type: application/json

{
  *"id": 1,
  "name": "John",
  "surname": "Smith",
  "phone": "+987654321",
  "company_id": 1,
  "passport": {
    "type": "ID",
    "number": "87654321"
  },
  "department": {
    "name": "Finance",
    "phone": "+123456789"
  }
}
```

**Ответ**:
```json
{
  "success": "Success"
}
```

---

## Тестирование

- Для тестирования REST API был использован **Postman**.
- Для тестирования gRPC можно использовать **evans**, **grpcurl** или плагины в IDE.

---

## Контейнеризация

Проект полностью контейнеризирован и включает следующие компоненты:

- **employee-service**: основной сервис для работы с сотрудниками.
- **api-gateway**: REST API Gateway.
- **postgres**: база данных.

## Структура проекта

- `/employee-service` — микросервис для работы с сотрудниками.
- `/api-gateway` — REST Gateway.
- `Dockerfile` — файлы для сборки контейнеров.
- `docker-compose.yml` — конфигурация Docker Compose.
