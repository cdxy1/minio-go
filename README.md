# Distributed File Storage System

Микросервисная система для хранения и управления файлами с использованием gRPC, REST API, Kafka и MinIO.

## Описание

Проект представляет собой распределенную систему хранения файлов, состоящую из трех основных сервисов:

- **File Service** - управление загрузкой и скачиванием файлов через MinIO
- **Metadata Service** - хранение метаданных файлов в PostgreSQL
- **Gateway** - REST API Gateway с документацией Swagger

## Архитектура

```
┌─────────────┐
│   Gateway   │ (REST API + Swagger)
│    :8080    │
└──────┬──────┘
       │
       ├─────┐
       │     │
┌──────▼──┐  │  ┌──────────────┐
│  File   │  │  │   Metadata   │
│ Service │  │  │   Service    │
│ :50051  │  └─▶│   :50052     │
└────┬────┘     └──────┬───────┘
     │                 │
     ▼                 ▼
┌─────────┐      ┌──────────┐
│  MinIO  │      │PostgreSQL│
│ :9000   │      │  :5432   │
└─────────┘      └──────────┘
     │                 │
     └────────┬────────┘
              ▼
         ┌─────────┐
         │  Kafka  │
         │  :9092  │
         └─────────┘
```

## Технологии

- **Go 1.25** - основной язык программирования
- **gRPC** - межсервисное взаимодействие
- **Gin** - REST API фреймворк
- **PostgreSQL** - база данных для метаданных
- **MinIO** - объектное хранилище
- **Kafka** - event-driven архитектура
- **Docker Compose** - оркестрация сервисов
- **Swagger** - документация API

## Быстрый старт

### Требования

- Go 1.25+
- Docker & Docker Compose
- Make

### Запуск проекта

1. **Клонирование репозитория:**
```bash
git clone https://github.com/cdxy1/minio-go.git
cd minio-go
```

2. **Сборка и запуск всех сервисов:**
```bash
make up
```

Эта команда:
- Соберет и запустит все сервисы через Docker Compose
- Инициализирует Kafka
- Поднимет PostgreSQL и MinIO

3. **Проверка работы:**
- Gateway API: http://localhost:8080
- Swagger UI: http://localhost:8080/swagger/index.html
- MinIO Console: http://localhost:9001 (admin123/admin123)
- PostgreSQL: localhost:5432

### Остановка

```bash
make down
```

## Структура проекта

```
.
├── cmd/                    # Точки входа сервисов
│   ├── file_service/      # File Service
│   ├── metadata_service/  # Metadata Service
│   └── gateway/           # Gateway Service
├── internal/
│   ├── app/               # Логика приложений
│   ├── config/            # Конфигурация
│   ├── entity/            # Доменные сущности
│   ├── grpc/              # gRPC handlers
│   ├── infra/             # Инфраструктура (Kafka, HTTP клиенты)
│   ├── repo/              # Репозитории
│   ├── routes/            # HTTP роуты
│   ├── service/           # Бизнес-логика
│   └── storage/           # MinIO и PostgreSQL подключения
├── api/
│   ├── proto/             # gRPC proto файлы
│   └── rest/              # Swagger документация
├── config/                # Конфигурационные файлы
├── dockerfiles/           # Docker файлы для сервисов
└── scripts/               # Вспомогательные скрипты
```

## Конфигурация

Конфигурация находится в директории `config/`:

- `config-dev.yaml` - конфигурация для разработки
- `config-local.yaml` - конфигурация для локального запуска

Основные параметры:
```yaml
postgres:
  host: postgres
  port: 5432
  user: admin
  password: 1234
  database: storage

minio:
  endpoint: minio:9000
  user: admin123
  password: admin123
  bucket: "vedro"

kafka:
  host: kafka:9092
  group: metadata-group
  topic: metadata-topic
```

## Разработка

### Сборка проекта

```bash
make build
```

Собирает все сервисы в директорию `bin/`:
- `bin/file`
- `bin/metadata`
- `bin/gateway`

### Генерация Proto файлов

```bash
make proto
```

Генерирует Go код из `.proto` файлов для gRPC сервисов.

### Проверка кода

```bash
make fmt  # Форматирование
make vet  # Статический анализ
```

### Запуск отдельных сервисов

```bash
# File Service
./bin/file

# Metadata Service
./bin/metadata

# Gateway
./bin/gateway
```

## API

### REST API (Gateway)

- `GET /swagger/*` - Swagger документация
- `POST /files/upload` - Загрузка файла
- `GET /files/download/:id` - Скачивание файла
- `GET /metadata` - Получить все метаданные
- `GET /metadata/:id` - Получить метаданные по ID

### gRPC API

#### File Service (:50051)

```protobuf
service FileService {
    rpc UploadFile(UploadFileRequest) returns (UploadFileResponse);
    rpc DownloadFile(DownloadFileRequest) returns (DownloadFileResponse);
}
```

#### Metadata Service (:50052)

```protobuf
service MetadataService {
    rpc GetAll(Empty) returns (FilesMetadataResponse);
    rpc GetById(FileMetadataRequest) returns (FileMetadataResponse);
}
```

## База данных

### PostgreSQL

База данных `storage` содержит таблицу метаданных файлов.

### MinIO

Объектное хранилище с бакетом `vedro` для хранения файлов.

## Kafka

Kafka используется для асинхронной обработки событий между сервисами.
- Topic: `metadata-topic`
- Group: `metadata-group`

## Docker

Все сервисы запускаются в Docker контейнерах:

```bash
# Поднятие всех сервисов
docker compose up -d

# Логи сервисов
docker compose logs -f gateway
docker compose logs -f file_service
docker compose logs -f metadata_service

# Остановка
docker compose down
```

## Лицензия

См. файл [LICENSE](LICENSE)

## Авторы

- [cdxy1](https://github.com/cdxy1)

---

## Troubleshooting

### Проблемы с портами

Убедитесь, что порты 8080, 5432, 9000, 9001, 9092 свободны.

### Kafka не запускается

```bash
docker exec kafka sh /kafka-init.sh
```

### Очистка данных

```bash
# Остановка и удаление volume
docker compose down -v
```
