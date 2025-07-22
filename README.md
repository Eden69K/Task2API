# API-шлюз для медицинских прогнозов

## Оглавление
* [Описание проекта](#описание-заголовка)
* [Техническая реализация](#техническая-реализация)
* [Функционал](#функционал)
* [Установка](#установка)
* [Запуск](#запуск)
* [Тестирование](#тестирование)

## Описание проекта
Данный API-шлюз выступает промежуточным слоем между клиентскими приложениями и сервисами машинного обучения, предоставляющими медицинские прогнозы. 

## Техническая реализация
### Структура проекта
```
project/
├── cmd/
│   └── app/
│       └── main.go         # Точка входа
├── config/
│   └── config.go           # Конфигурация
├── internal/
│   ├── handler/            # Обработчики HTTP
│   ├── models/             # Модели данных
│   ├── interfaces/         # Интерфейсы сервисов
│   ├── services/           # Бизнес-логика
│   ├── middleware/         # Промежуточное ПО
│   └── routes/             # Маршрутизация
├── pkg/
│   └── logger/             # Логирование
├── release/                # Исполняймые бинарники
├── scripts                 # Скрипты 
├── go.mod                  # Зависимости
├── go.sum                  # Зависимости
├── Dockerfile              # Конфигурация образа
├── docker-compose.yml      # Оркестрация сервисов
├── .env                    # Файл окружения
└── config.yml              # Конфигурационный файл
```
### Описание параметров
|Параметр|Тип|Описание|
|-|-|-|
|age|int|Возраст пациента|
|gender|int|Пол пациента(0 - мужской, 1 - женский)|
|rdw|float|Ширина распределения эритроцитов (%)|
|wbc|float|Лейкоциты (10^9/л)|
|rbc|float|Эритроциты (10^12/л)|
|hgb|float|Гемоглобин (г/л)|
|hct|float|Гематокрит (%)|
|mcv|float|Средний объем эритроцита (фл)|
|mch|float|Среднее содержание гемоглобина в эритроците (пг)|
|mchc|float|Cредняя концентрация гемоглобина в эритроците (г/дл)|
|plt|float|Тромбоциты (10^9/л)|
|neu|float|Нейтрофилы (%)|
|eos|float|Эозинофилы (%)|
|bas|float|Базофилы (%)|
|lym|float|Лимфоциты (%)|
|mon|float|Моноциты (%)|
|soe|float|Скорость оседания эритроцитов (мм/ч)|
|chol|float|Холестерин (ммоль/л)|
|glu|float|Глюкоза (ммоль/л)|
|hdl|float|Липопротеины высокой плотности (ммоль/л)|
|tg|float|Триглицериды (ммоль/л)|
|cpr|float|C-реактивный белок (мг/л)|

## Функционал
### Доступные эндпоинты
|Эндпоинт|Параметры|Описание|
|-|-|-|
|/health|---|Проверка работоспособности сервиса|
|/predict/hba1c|"age", "gender", "rdw", "wbc", "rbc", "hgb", "hct", "mcv", "mch", "mchc", "plt", "neu", "eos", "bas", "lym", "mon", "soe", "chol", "glu"|Прогноз HbA1c(Гликированный гемоглобин)|
|/predict/ldl|"age","gender","rdw","wbc","rbc","hgb","hct","mcv","mch","mchc","plt","neu","eos","bas","lym","mon","soe","chol","glu"|Прогноз LDL(Липопротеины низкой плотности)|
|/predict/ldll|"age","gender","chol","hdl","tg"|Прогноз LDLL(производный показатель LDL)|
|/predict/ferr|"age","gender","rdw","wbc","rbc","hgb","hct","mcv","mch","mchc","plt","neu","eos","bas","lym","mon","soe","crp"|Прогноз FERR(Ферритин)|
|/predict/tg|"age","gender","rdw","wbc","rbc","hgb","hct","mcv","mch","mchc", "plt","neu","eos","bas","lym","mon","soe","chol","glu"|Прогноз TG(Триглицериды)|
|/predict/hdl|"age","gender","rdw","wbc","rbc","hgb","hct","mcv","mch","mchc","plt","neu","eos","bas","lym","mon","soe","chol","glu"|Прогноз HDL(Холестерин липопротеинов высокой плотности)|

## Установка
```
git clone https://github.com/Eden69K/Task2API.git
cd Task2API
```

## Запуск
### Запуск через docker
```
docker-compose up -d --build
```
### Бинарные сборки
#### Windows
**Запуск**
```
./release/windows/api-gateway.exe
```
**Создание исполняемого файла(если он отсуствует)**
```
./scripts/buildwin.ps1
```
#### Linux
**Запуск**
```
chmod +x ./release/linux/api-gateway
./release/linux/api-gateway
```
**Создание исполняемого файла(если он отсуствует)**
```
chmod +x ./scripts/buildlin.sh
./scripts/buildlin.sh
```
## Тестирование
### Тестирование через Postman
**Пример запроса**
```
GET http://localhost:8080/predict/hba1c
Headers:
  Authorization: Bearer {auth_token}
```

**Health check**
```
GET http://localhost:8080/health
Headers:
  Authorization: Bearer {auth_token}
```
