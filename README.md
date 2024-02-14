# Local Setup

## Create a postgres db with docker

```bash
$ docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

$ docker exec -it postgres15 createdb --username=admin --owner=admin stori_db
```

## Drop database if needed

```bash
$ docker exec -it postgres15 dropdb --username=admin stori_db
```

## Install golang-migrate

_docs:_ https://github.com/golang-migrate/migrate?tab=readme-ov-file

**mac:**

```bash
$ brew install golang-migrate
```

**windows**

```bash
$ scoop install migrate
```

**linux**

```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

## Migrate tables with golang-migrate

```bash
$ migrate -path db/migrations -database "postgresql://admin:secret@127.0.0.1:5432/stori_db?sslmode=disable" --verbose up
```

## Install sqlc

_Just in case we want to continue with the development, we need sqlc to generate code_

```bash
$ brew install sqlc
```

## Create a .env file

```bash
PORT=8000
DB_HOST="localhost"
DB_USER="admin"
DB_PASSWORD="secret"
DB_NAME="stori_db"
IS_LOCAL=true

EMAIL_HOST=#mailtrap host
EMAIL_PORT=#mailtrap port
EMAIL_USER=#mailtrap user
EMAIL_PASSWORD=#mailtrap pass

FRONT_URL="http:localhost:3000"

```

## Run the program

```bash
$ go run .
```

## Postman collection to test locally

```bash
https://api.postman.com/collections/6961632-baa5b9fa-4e5e-4242-88a7-c209ae4f65d3?access_key=PMAT-01HJQ4N7B40ZM5PAPSRFYJYZ2Z
```

## Test files

_test files are included in `transfers_file` directory_

## Front

you can use this project with the front end: https://github.com/tecolotedev/stori_front
