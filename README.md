#  golang-hashtag

This project is created with gofiber v2.

## Installation

1) Clone
```bash
git clone https://github.com/pd-pranay/golang-hashtag.git
```
2) Start database with docker
```bash
docker-compose up
```
3) Run migrations 
```bash
cd db/migrations
goose postgres "postgresql://localhost:7777/hashtag?user=hashtag&password=hashtag&sslmode=disable" status
```
3) Run project 
```bash
go run main.go
```

## Packages used

```
github.com/gofiber/fiber/v2
github.com/lib/pq
github.com/google/uuid
https://github.com/kyleconroy/sqlc

```
