# Data-Aggregation-service
REST API сервис для агрегации данных об онлайн-подписках пользователей

## Описание API
Сервис включает функционал создания, просмотра, обновления и удаления подписок, а также расчет их суммарной стоимости.

## Запуск
`docker-compose -f local.docker-compose.yaml up -d`

> [!NOTE]
> Сервис поддерживает только локальное развертывание, файл конфигурации local.yaml, а так же файлы с переменными окружения app-local.env и db-local.env представлены в репозитории и являются локальными образцами.
> 
> Вся локальная конфигурация подтягивается автоматически командой выше.

## Базовый URL
`http://localhost:8080/api/v1`

## Swagger
Swagger (OpenAPI) документация представлена в /api

## Примеры запросов
### 1. POST /subscriptions (создание подписки)
```json
curl -X POST "http://localhost:8080/api/v1/subscriptions" \
  -H "Content-Type: application/json" \
  -d '{
    "service_name": "Yandex Plus",
    "price": 200,
    "user_id": "eb8a32db-139f-4e33-b172-39810efcc487",
    "start_date": "08-2025",
    "end_date": "09-2025"
  }'
```
### 2. GET /subscriptions (получение списка подписок с фильтрацией)
```json
curl -X GET "http://localhost:8080/api/v1/subscriptions?user_id=eb8a32db-139f-4e33-b172-39810efcc487&service_name=Yandex+Plus"
```
### 3. GET /subscriptions/{subId} (получение конкретной подписки по subId)
```json
curl -X GET "http://localhost:8080/api/v1/subscriptions/248d128a-91a6-454e-b01f-c85ee5ca0471"
```
### 4. PATCH /subscriptions/{subId} (частичное обновление подписки по subId)
> [!NOTE]
> Поддерживается обновление цены и/или даты окончания подписки.
```json
curl -X PATCH "http://localhost:8080/api/v1/subscriptions/248d128a-91a6-454e-b01f-c85ee5ca0471" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 700,
    "end_date": "04-2025"
  }'
```
### 5. DELETE /subscriptions/{subId} (удаление подписки по subId)
```json
curl -X DELETE "http://localhost:8080/api/v1/subscriptions/248d128a-91a6-454e-b01f-c85ee5ca0471"
```
### 6. GET /subscriptions/cost/total (получение суммарной стоимости подписок с фильтрацией)
```json
curl -X GET "http://localhost:8080/api/v1/subscriptions/cost/total?from=01-2025&to=12-2025&user_id=eb8a32db-139f-4e33-b172-39810efcc487"
```