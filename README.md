## Установка и запуск проекта

### Клонирование репозитория
Для начала работы склонируйте репозиторий в удобную Вам директорию:
```bash
git clone https://github.com/kstsm/task-api
```
### Настройка переменных окружения
Создайте `.env` файл, скопировав в него значения из `.env.example`, и укажите необходимые параметры.
## Запуск приложения

#### Способ 1: Прямой запуск
```bash
go run main.go
```
#### Способ 2: Через Makefile
```bash
make start
```

После запуска сервер будет доступен по адресу: `http://localhost:8080`

## API Endpoints

### Базовый URL
```
http://localhost:8080/api/v1/task
```

### 1. Создание задачи

**POST** `http://localhost:8080/v1/task`

Создает новую задачу и запускает её асинхронную обработку.

### 2. Получение задачи

**GET** `http://localhost:8080/v1/task/{id}`

Получает информацию о задаче по её UUID.

### 3. Удаление задачи

**DELETE** `http://localhost:8080/v1/task/{id}`

Удаляет задачу по её UUID.


