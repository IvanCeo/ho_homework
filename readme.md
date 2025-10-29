# Домашнее задание по golang

---

Скопировать репозиторий:
```bash
git clone https://github.com/IvanCeo/ho_homework/tree/master
```

## Сервис Ledger

Для запуска теста работы сервиса Ledger запустите команды
```bash
cd hmwk2/ledger/cmd
make test
```

## Сервис Gateway

В сервисе доступны 4 эндпоинта: создание транзакции, создание бюджета, список транзакций, список бюджетов.
Для запуска теста работы сервиса Gateway запустите команды
```bash
cd hmwk2/gateway/cmd/server
go run main.go
```

1. Добавить несколько бюджетов
```bash
curl -X POST http://127.0.0.1:8080/api/budgets \
  -H "Content-Type: application/json" \
  -d '{
    "category": "Food",
    "limit": "3000",
    "Period": "1291309712"
  }'
```
```bash
curl -X POST http://127.0.0.1:8080/api/budgets \
  -H "Content-Type: application/json" \
  -d '{
    "category": "Transport",
    "limit": "2000",
    "Period": "1291309712"
  }'
```

2. Добавить несколько транзакций
```bash
curl -X POST http://127.0.0.1:8080/api/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "amount": "1200",
    "category": "Food",
    "description": "Dinner in cafe",
    "date": "2025-04-01T12:15:23.52Z"
  }'
```
```bash
curl -X POST http://127.0.0.1:8080/api/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "amount": "200",
    "category": "Transport",
    "description": "Taxi to cafe",
    "date": "2025-04-01T12:10:23.52Z"
  }'
```
```bash
curl -X POST http://127.0.0.1:8080/api/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "amount": "2000",
    "category": "Food",
    "description": "Breakfast on Duty",
    "date": "2025-04-03T22:23:07.52Z"
  }'
```

3. Список бюджетов
```bash
curl http://127.0.0.1:8080/api/budgets

```

4. Список транзакций
```bash
curl http://127.0.0.1:8080/api/transactions
```