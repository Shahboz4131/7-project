#!/usr/bin/env sh
set -e

>&2 echo "Running migration ..."
# migrate -path=./migrations -database=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:5432/$POSTGRES_DB?sslmode=disable up
migrate -path=./migrations -database=postgres://postgres:postgres@localhost:5432/task?sslmode=disable up

tail -f /dev/null