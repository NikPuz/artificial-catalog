При разработке использовалось:
- Среда разработки Visual Studio Code
- библиотеки gorilla/mux, viper, zap, goose
- Система управления базами данных MySQL
- Локальный веб-сервер Open Server Panel
- Система контроля версий GitHub
- Система контейнеризации Docker

Сервис поддерживает следующие запросы:
- GET /
- GET /dai
- GET /image/{imageName}

Запуск:
- Установленный Docker https://www.docker.com/products/docker-desktop/
- `$ docker compose --env-file ./config/app.env up` в папке проекта
