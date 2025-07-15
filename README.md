# API-шлюз для медицинских прогнозов

## Описание проекта
Данный API-шлюз выступает промежуточным слоем между клиентскими приложениями и сервисами машинного обучения, предоставляющими медицинские прогнозы. 

## Техническая реализация
### Структура проекта
```
project/
├── cmd/
│   └── app/
│       └── main.go          # Точка входа
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
├── go.mod                  # Зависимости
├── go.sum                  # Зависимости
├── README.md               # Документация
└── config.yml              # Конфигурационный файл
```

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
