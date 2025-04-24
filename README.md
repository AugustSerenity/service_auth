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

### Тестовые данные для создания пары Access, Refresh токенов

#### Тестовый запрос из командной строки 
```sh
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

### Тестовые данные для Refresh операции Access, Refresh токенов
```sh
curl -X POST http://localhost:8080/auth-refresh \                                      
  -H "Content-Type: application/json" \
  -d '{
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzZTQ1NjctZTg5Yi0xMmQzLWE0NTYtNDI2NjE0MTc0MDAwIiwiaXAiOiIxOTIuMTY4LjY1LjEiLCJleHAiOjE3NDU1MjM4MjksImlhdCI6MTc0NTUyMDIyOSwianRpIjoiMTFhMzM1MjFhZjM2M2VlYSJ9.1j49U9Fl3X4SJBHAunD6l0DVs-8s4x3d4nOXGvOpo9jaMOOMzNCIAsAIDt_3uotVnKR0286cu2UPcQD1TzSNwg",
    "refresh_token": "MGEyODk0OGM3ZWYwZGYwOWJmNWM0OTRmNzUxMTJiOWE="
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
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzZTQ1NjctZTg5Yi0xMmQzLWE0NTYtNDI2NjE0MTc0MDAwIiwiaXAiOiIxOTIuMTY4LjY1LjEiLCJleHAiOjE3NDU1MjM4MjksImlhdCI6MTc0NTUyMDIyOSwianRpIjoiMTFhMzM1MjFhZjM2M2VlYSJ9.1j49U9Fl3X4SJBHAunD6l0DVs-8s4x3d4nOXGvOpo9jaMOOMzNCIAsAIDt_3uotVnKR0286cu2UPcQD1TzSNwg",
    "refresh_token": "MGEyODk0OGM3ZWYwZGYwOWJmNWM0OTRmNzUxMTJiOWE="
  }
}
```
