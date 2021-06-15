# webservice

## Запуск webservice

Команды для запуска

docker build -t webservice .

docker run -p 127.0.0.1:8080:8080 webservice

## Запросы

url ```http://localhost:8080```

### Добавление нового пользователя

```POST``` /user?name=Hahan&birth_date=2000-12-20

response:

      200 - OK
      400 - Bad Request
      500 - Internal Server Error

### Получить данные по пользователю

```GET``` /user?id=6

response:
      
      200 - OK
```
{
"id": 6,
"name": "Hahan",
"birth_date": "2000-12-20T00:00:00Z"
}
```

      400 - Bad Request
      404 - Not Found
      500 - Internal Server Error

### Удаление пользователя

```DELETE``` /user?id=6

response:

      200 - OK
      400 - Bad Request
      500 - Internal Server Error

### Обновление пользователя
   
```PUT``` /user?id=6

body:
```
{
"name": "Hahan",
"birth_date": "2000-12-20"
}
```

response:

      200 - OK
      400 - Bad Request
      404 - Not Found
      500 - Internal Server Error