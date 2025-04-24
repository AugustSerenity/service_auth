# service_auth
service_auth

# Тестовое задание для отбора BackDev 
## auth_service

## Описание
Написать часть сервиса аутентификации.

Два REST маршрута:
- Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса
- Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов

## Начало работы
### Установка
Клонирование репозитория
```sh
git clone https://github.com/AugustSerenity/service_auth
```
### Запуск сервиса
Запускаем контейнер с помощью Makefile
```sh
make run
```

### Тестовый запрос для для создания пары Access, Refresh токенов

```bash
curl -X POST "http://localhost:8080/auth-token?id=123e4567-e89b-12d3-a456-426614174000"
```

##### JSON

```json
{
  "method": "POST",
  "url": "http://localhost:8080/auth-token",
  "params": {
    "id": "123e4567-e89b-12d3-a456-426614174000"
  },
  "headers": {
    "Content-Type": "application/json"
  }
}
```

### Тестовый запрос для операции refresh

Для выполнения запроса на обновление токенов, используйте следующий `curl` запрос:

```bash
curl -X POST "http://localhost:8080/auth-refresh" \
  -H "Content-Type: application/json" \
  -d '{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzZTQ1NjctZTg5Yi0xMmQzLWE0NTYtNDI2NjE0MTc0MDAwIiwiaXAiOiIxOTIuMTY4LjY1LjEiLCJleHAiOjE3NDU1MjYxMjgsImlhdCI6MTc0NTUyMjUyOCwianRpIjoiYTMwMTI0YmVkOTg4ZmViOSJ9.GJwHIUiMxEqCdHg_Km2wg6Oq7_sFLxmk9llx9INakmTfklzMIcYpOcdLXpOI65acwIC0WqWP_adV892Bmb8h5Q",
    "refresh_token": "MTFmZmZkNjdkODk2ZTMzZmUxYTg1Yzk1OTMxNDNiMGY="
  }'
```

##### JSON
```json
{
  "method": "POST",
  "url": "http://localhost:8080/auth-refresh",
  "headers": {
    "Content-Type": "application/json"
  },
  "body": {
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzZTQ1NjctZTg5Yi0xMmQzLWE0NTYtNDI2NjE0MTc0MDAwIiwiaXAiOiIxOTIuMTY4LjY1LjEiLCJleHAiOjE3NDU1MjYxMjgsImlhdCI6MTc0NTUyMjUyOCwianRpIjoiYTMwMTI0YmVkOTg4ZmViOSJ9.GJwHIUiMxEqCdHg_Km2wg6Oq7_sFLxmk9llx9INakmTfklzMIcYpOcdLXpOI65acwIC0WqWP_adV892Bmb8h5Q",
    "refresh_token": "MTFmZmZkNjdkODk2ZTMzZmUxYTg1Yzk1OTMxNDNiMGY="
  }
}
```
