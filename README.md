# Интрукция по запуску
1. Запустить контейнер с базой данных<br>
`docker compose up db`
2. Запустить контейнер с сервером<br>
`docker compose up server`

# Эксплуатация
OpenAPI документация лежит в /api

## Основное задание
1. Резервирование товаров по массиву уникальных кодов<br>
`curl --location 'localhost:8080/warehouse/items/reserve' \
   --header 'Content-Type: application/json' \
   --data '{
    "codes": [
        "cdbnth1",
        "eelb2"
    ]
}'`
2. Освобождение резерва товаров по массиву уникальных кодов<br>
`curl --location 'localhost:8080/warehouse/items/release' \
   --header 'Content-Type: application/json' \
   --data '{
    "codes": [
        "cdbnth1",
        "eelb2"
    ]
}'`
3. Получение товаров на складе по id склада<br>
`curl --location 'localhost:8080/warehouse/items/left' \
   --header 'Content-Type: application/json' \
   --data '{
    "id": 1
}'`

Проверять правильность резервирования и освобождения можно путем совершения запросов 3 -> 1 -> 3 -> 2 -> 3 и так далее с изменением id склада и кодов товаров

## Дополнительное задание
4. Получение товаров со всех складов по id товара<br>
`curl --location 'localhost:8080/warehouse/item' \
   --header 'Content-Type: application/json' \
   --data '{
    "id": 1
}'`

# База данных
1. Таблица item
![item](https://github.com/teatou/lamoda/raw/main/api/pics/item_sql.png)

2. Таблица warehouse
![item](https://github.com/teatou/lamoda/raw/main/api/pics/warehouse_sql.png)

3. Таблица warehouse_item
![item](https://github.com/teatou/lamoda/raw/main/api/pics/warehouse_item_sql.png)