# Сервис предоставляющий информацию о заказе

Демонстрационный сервис с простейшим интерфейсом для получения данных по некоторому uid-заказа.

## Требования ⚙️
Для запуска этого сервиса вас потребуется следующее:
- Docker
- Создать в корневой директории файл .env и заполнить по шаблону:
```
#POSTGRES ENVIRONMENTS
PGUSER=user
PGPASSWORD=password
PGHOST=postgres
PGPORT=5432
PGDATABASE=db
PGSSLMODE=disable

#SERVER ENVIRONMENT
HTTP_PORT=8080

#NATS-STREAMING
CLUSTER_ID=test-cluster
CLIENT_ID=test-client
LISTEN_CHANNEL=orders
LISTEN_URL=http://nats-streaming:4222/
```

#### Миграции для базы данных находятся в директории [/migration](./migrations)

## Запуск 🔧

Для запуска выполните в терминале команду ```make compose-up```, после чего сервер будет запущен на localhost на указанном
вами порту.
Для остановки сервера нужно прописать команду ```make compose-down```

Запуск тестов командой `make test`, запуск тестов с покрытием `make cover` и для получения отчёта в html формате `make cover-html`

## Интерфейс 🌐
После успешного запуска и перехода по пути ```http://localhost:8080/order``` (если указан тот же порт, что и в шаблоне) 
у вас откроется страница, где вы можете ввести номер заказа и получить информацию по заказу.

## Публикация в канал nats-streaming
В директории есть скрипт [publisher.go](./publisher.go), с помощью него можно реализовать запись в nats-streaming канал
(считывает данные о заказе из файла [model.json](./model.json))