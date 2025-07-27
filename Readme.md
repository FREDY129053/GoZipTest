## 📋 Описание задачи

Сервис предоставляет API для скачивания файлов по ссылкам из интернета, упаковки их в .zip архив и предоставления этого архива пользователю.

Ограничения:
- В архиве может быть максимум 3 файла.

- Допустимые расширения: .jpg, .pdf.

- Одновременно может выполняться не более 3 задач по созданию архива.

- При недоступности одного или нескольких ресурсов, в архив добавляются только успешно скачанные файлы. Информация об ошибках возвращается в ответе.


## 🚀 Запуск проекта

- Проект запускается с помощью **Makefile**.

    ```bash
    make run-api
    ```
- Проект можно запустить самостоятельно.
    ```bash
    cd app && go run cmd/main.go
    ```

После этого сервис будет доступен по адресу:
http://localhost:8080


## ⚙️ Структура API
### Основные возможности:

1. Создание задачи на архивирование.

2. Добавление ссылок к задаче.

3. Получение статуса задачи и ссылки на архив.

4. Скачивание архива по готовой задаче.

### Swagger документация:
Доступна после запуска проекта:
http://localhost:8080/docs/index.html


## Структура проекта
```bash
app/
├── archives/               # Сюда сохраняются готовые zip-архивы
├── cmd/main.go             # Точка входа
├── docs/                   # Сгенерированная Swagger-документация
├── internal/
│   ├── api/handlers/       # Обработчики HTTP-запросов
│   ├── api/routers/        # Настройка роутинга
│   ├── database/           # Псевдо-БД (in-memory)
│   ├── models/             # Модели
│   ├── repository/         # Работа с задачами и архивами
│   ├── schemas/            # Валидация и Swagger-схемы
│   └── service/            # Логика архивации
├── pkg/helpers/            # Утилиты (архиватор и т.п.)
├── test.http               # Набор тестовых HTTP-запросов
└── Makefile
```

## 🛠 Используемые технологии

- [Go](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://go.dev/&ved=2ahUKEwie0aCr9dyOAxUS_AIHHT4MOqYQFnoECBoQAQ&usg=AOvVaw2GopNGe_pMsOKwFeS7coZ3) (v1.24+)

- [Gin](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://gin-gonic.com/&ved=2ahUKEwj7-fqy9dyOAxW42wIHHf9gAf4QFnoECAoQAQ&usg=AOvVaw3NpIRVflxdDy9iVoKMyoRA) — HTTP-фреймворк

- [swaggo/swag](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://github.com/swaggo/swag&ved=2ahUKEwiXuve79dyOAxW_8gIHHcsqD8wQFnoECAoQAQ&usg=AOvVaw1mZ7Ln626ByV8c3M7g8hbY) — автогенерация Swagger-документации

## ✅ Принятые архитектурные решения и паттерны

- Чистая архитектура: внутренние пакеты разделены по слоям (handler → service → repository).

- In-Memory хранилище для хранения задач (никакие внешние базы данных не используются).