# Snake PVP


## Стек
- **Backend**: Go + Gin + Centrifuge (WebSocket)
- **Frontend**: Vue 3

## Запуск в dev окружении
1. docker compose up
2. http://localhost:5173

## Билд
1. ```docker compose exec -it node bun run build``` - Билд фронта
2. ```docker compose exec -it app go build cmd/snake/main.go ``` - Билд бинарного файла