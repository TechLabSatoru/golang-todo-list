## 環境構築

- PostgreSQLの場合

```shell
$ docker compose up -d

# PostgreSQLの場合
$ docker compose exec -it postgres bash

$ psql -h localhost -U postgres -d postgres

postgres=# 

$ docker compose down --rmi all --volumes --remove-orphans
```

- Golangの場合

```shell
$ docker compose up -d

# Golangの場合
$ docker compose exec -it golang 

$ cd golang-todo-list
$ go run main.go

$ docker compose down --rmi all --volumes --remove-orphans
```
