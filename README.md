# Интрукция по запуску
1. Запустить контейнер с базой данных<br>
`docker compose up db`
2. Запустить контейнер с сервером<br>
`docker compose up server`

# Эксплуатация
OpenAPI документация лежит в /api

## Основное задание
1. Резервирование товаров<br>
`curl --location 'localhost:8080/warehouse/items/reserve' \
   --header 'Content-Type: application/json' \
   --data '{
    "codes": [
        "cdbnth1",
        "eelb2"
    ]
}'`
2. Освобождение резерва товаров<br>
`curl --location 'localhost:8080/user/remove' \
   --header 'Content-Type: application/json' \
   --data '{
    "codes": [
        "cdbnth1",
        "eelb2"
    ]
}'`
3. Получение товаров на складе<br>
`curl --location 'localhost:8080/warehouse/items/left' \
   --header 'Content-Type: application/json' \
   --data '{
    "id": 1
}'`

Проверять правильность резервации и освобождения можно путем совершения запросов 3 -> 1 -> 3 -> 2 -> 3 и так далее с изменением id склада и кодов товаров

## Дополнительное задание 3
4. Создание сегмента с изначальными пользователями<br>
`curl --location 'localhost:8080/segment/add' \
   --header 'Content-Type: application/json' \
   --data '{
   "slug": "segment_4",
   "percent": 50
   }'`