# Camparser
Софт позволяет выполнить любой запрос shodan, получить результаты,сохранить их в базе,выводить данные на карту в соответствии с координатами.

# Запуск

Оставил свой ключ шодана для теста в docker-compose.yml

Запуск осуществляется при помощи docker compose.
Docker-compose файл включает в себя базу данных,pgadmin,само приложение.

```
docker-compose up -d
```

Веб интерфейс будет доступен по адресу http://127.0.0.1:8081/.


# Описание архитектуры
Данное приложение построено опираясь на чистую архитектуру Uncle Bob.
```
.
├── cmd //main файл 
├── configs //конфиги
├── internal //внутренние пакеты
│   ├── handler //обработчики веба
│   ├── repository //репозиторий
│   ├── server //сервер
│   └── service //бизнес логика
├── pkg //Общие пакеты
│   └── shodan //пакет для работы с шоданом
├── public  //веб
├── schema  //миграции бд
├── screenshots //скрины
├── types.go //типы уровня domain
└── wait-for-postgres.sh //скрипт для ожидания запуска бд в docker compose
```

# Описание приложения

Связь фронтэнда с бэкэндом осуществляется через api.
Скриншот с postman:
![Скриншот с postman](/screenshots/postman.jpg "Скриншот с postman").

Парсинг камер осуществляется в го рутинах, для мониторинга за процессом данные о задачах вносятся в бд.
Если к базе данных небыли применены миграции, то приложение само создаст нужные таблицы.


# Описание контейнеризации

Сборка производится в 2 этапа, для сокращения веса готового образа. Такие данные как пароли и токены передаются через переменные окружения. Также образ включает в себя файл wait-for-postgres.sh. Этот файл запустит приложение только после отклика от базы данных. При запуске через docker compose приложение запустится быстрее чем база данных и выпадет в панику.


