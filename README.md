# Wallet

## Install tools

* [Docker desktop](https://www.docker.com/products/docker-desktop)
* [Golang](https://golang.org/)
* [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* [Sqlc](https://github.com/kyleconroy/sqlc#installation)
* [Gomock](https://github.com/golang/mock)

## How to generate code

* Generate SQL CRUD with sqlc:

```
make sqlc
```

* Generate DB mock with gomock:
```
make mock
```

## How to run

0) Setup config parameters in app.env.

1) Generate SQL CRUD with sqlc:
```
make sqlc
```

2) Run docker-compose:

```
docker-compose up
```
3) Make migrations:
```
make migrateup
```
4) Use curl examples to test api. Create new user account first, then perform operations.


* Run test:
```
make createtestdb
make migratetestup
make test
```

## Curl examples

Create account
```
curl -d '{"owner":"test user 1", "currency":"USD"}' -H "Content-Type: application/json" -X POST http://wallet.app.loc:8080/accounts
```

Deposit Operation
```
curl -d '{"user_id":1, "amount":200, "currency":"USD", "description": "add some money"}' -H "Content-Type: application/json" -X POST http://wallet.app.loc:8080/deposits
```

Withdrawal Operation

```
curl -d '{"user_id":1, "amount":20, "currency":"USD", "description": "withdrawal operation"}' -H "Content-Type: application/json" -X POST http://wallet.app.loc:8080/withdrawals
```

Get balance (GET request - /balance/{user_id})

```
curl -X GET http://wallet.app.loc:8080/balance/1
```

Get list of transactions 

```
curl -X GET "http://wallet.app.loc:8080/transactions?start_date=2021-01-01&stop_date=2021-07-19&page_id=1&page_size=10"
```


