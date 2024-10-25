set dotenv-load := true

# Define variables

DB_DIR := "db"
MIGRATIONS_DIR := "db/schema"

default:
    just --fmt --unstable
    just --list

# clean all build artifacts
clean:
    rm -rf bin

# build all services
build:
    go build  -o bin/pih-core main.go

# run all services
run:
    air -c .air.toml

# apply all db migrations
goose-up:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL up

# rollback last db migration
goose-down:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL down

# show the status of db migrations
goose-status:
    goose -dir {{ MIGRATIONS_DIR }} postgres $DB_URL status

# create a new db migration with name
goose-create name:
    goose -dir {{ MIGRATIONS_DIR }} create {{ name }} sql

# vendor all dependencies
vendor:
    go mod vendor

# run go mod tidy on all services
tidy:
    go mod tidy
