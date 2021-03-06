
NAME=User
FN=user

.PHONY: start
start:
	docker-compose up -d

.PHONY: build
build:
	docker-compose up -d --build

.PHONY: log
log:
	docker-compose logs

.PHONY: log
log:
	docker-compose logs

.PHONY: log-api
log-api:
	docker logs -f api

.PHONY: log-db
log-cron:
	docker logs -f db

.PHONY: kill
kill:
	docker-compose kill

.PHONY: down
down:
	docker-compose down --volumes

.PHONY: ps
ps:
	docker-compose ps

.PHONY: exec-api
exec-api:
	docker exec -it api /bin/sh

.PHONY: exec-db
exec-db:
	docker exec -it db bin/bash

.PHONY: test
test:
	docker exec api go test -v

.PHONY: test-auth
test-auth:
	docker exec api go test -v github.com/kichikawa/auth/... -count=1

.PHONY: test-infra
test-infra:
	docker exec api go test -v github.com/kichikawa/infra/... -count=1

.PHONY: restart
restart: kill start

.PHONY: schema-add
schema-add:
	docker exec api go run entgo.io/ent/cmd/ent init ${NAME}

.PHONY: schema-gen
schema-gen:
	docker exec api go generate ./...

.PHONY: gqlgen
gqlgen:
	docker exec api gqlgen