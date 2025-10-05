# CriticDealer

**CriticDealer** — сервис для анализа аварийности маршрутов на основе данных 2GIS и OpenWeather.

---

## 🚀 Возможности

- Получает маршруты от **2GIS Routing API**.
- Загружает погоду с **OpenWeather API**.
- Сопоставляет с историческими данными о ДТП.
- Рассчитывает вероятность критичности каждого манёвра.
- Поддерживает сезонный и погодный контекст анализа.

---

## 🧰 Технологии

- **Go (Golang)** — серверная логика
- **Gin** — фреймворк для REST API
- **PostgreSQL** — хранилище данных
- **2GIS Routing API** — маршрутизация
- **OpenWeather API** — погода
- **CRC32** — хэширование манёвров
- **ENUM PostgreSQL** — типы погоды и дней недели

---

## ⚙️ Конфигурация окружения

Создайте файл `.env` в корне проекта и укажите параметры:

```dotenv
DB_CONNECT_STRING="postgres://postgres:password@localhost:5432/criticdealer?sslmode=disable"
SERVER_PORT=":8080"
JWT_SECRET="supersecret"
LOG_LEVEL="info"
LOG_TO_CONSOLE="true"

WEATHER_KEY="YOUR_OPENWEATHER_API_KEY"
2GIS_KEY="YOUR_2GIS_API_KEY"
```

---

## 🧩 Структура проекта

```
internal/
├── api/                # Основные обработчики маршрутов
├── config/             # Конфигурации и логирование
├── models/
│   ├── db/             # Модели БД
│   └── routresponse/   # Структуры 2GIS и OpenWeather
├── repository/         # Работа с PostgreSQL
├── service/
│   ├── funcmonth/      # Определение месяца и типа дня
│   ├── ids/            # Генерация CRC32-хэшей манёвров
│   ├── math/           # Алгоритм расчёта критичности
│   └── weather/        # Запрос и парсинг погоды
└── cmd/main.go         # Точка входа
```

---

## 🧮 Формула критичности

```go
probability = (currentWeatherAccidents / (length * avgTraffic)) +
              (totalAccidents / (length * avgTraffic) * clearRatio)
```

Где:
- `length` — длина манёвра (в метрах)
- `avgTraffic` — средний поток
- `clearRatio` — доля аварий в ясную погоду

---

## 🗄️ Таблицы БД

### accident
| Поле | Тип | Описание |
|------|-----|-----------|
| id | serial | Первичный ключ |
| hash | bigint | CRC32-хэш манёвра |
| dtp_time | integer | Час аварии |
| month | integer | Месяц |
| traffic | integer | Интенсивность потока |
| day_type | enum | Weekday/Weekend |
| weather_id | int | Ссылка на таблицу погоды |

### weather
| id | serial | PK |
| weather_type | enum | Погодное состояние |

### global_accident_statistic
| dtp_count | int | Кол-во ДТП |
| dtp_koef | decimal | Средний коэффициент |
| region | varchar | Регион |

---

## 🧠 Пример запроса

```json
POST /api/v1/critical
{
  "points": [
    {"lon": 37.621365, "lat": 55.847874},
    {"lon": 37.615781, "lat": 55.850148, "type": "stop"}
  ],
  "transport": "driving",
  "output": "detailed",
  "locale": "ru",
  "alternative": 3
}
```

---

## 💾 Пример ответа

```json
{
  "status": "OK",
  "result": [
    {
      "id": "5467544109162600749",
      "maneuvers": [
        {
          "comment": "Поворот направо на ул. Берёзовая аллея",
          "critical_probability": 0.0023
        }
      ]
    }
  ]
}
```

---

## 🧪 Тестовые данные

Используйте следующий `.env.example`:

```dotenv
DB_CONNECT_STRING="postgres://postgres:197320@localhost:5432/2gis?sslmode=disable"
SERVER_PORT=":8080"
JWT_SECRET="supersecret"
LOG_LEVEL="info"
LOG_TO_CONSOLE="true"
WEATHER_KEY="7e8a6ed72cffae478833bb79e3f7e194"
2GIS_KEY="9abec63a-4211-4ea4-94ae-8a0d41c26d81"
```

---

## 📦 Запуск

```bash
go mod tidy
go run cmd/main.go
```

---

## 🌦️ Пример вызова OpenWeather API

```
https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=YOUR_WEATHER_KEY
```

---

## 🧱 Пример миграции

```sql
CREATE TYPE weather AS ENUM (
  'Clear','Clouds','Rain','Snow','Drizzle','Thunderstorm',
  'Mist','Fog','Haze','Ice','Freezing_Rain'
);

CREATE TYPE day_type AS ENUM ('Weekday','Weekend');
CREATE TABLE accident (
  id serial PRIMARY KEY,
  hash BIGINT UNIQUE NOT NULL,
  dtp_time integer,
  month integer,
  traffic integer,
  day_type day_type,
  weather_id integer REFERENCES weather(id)
);
```
