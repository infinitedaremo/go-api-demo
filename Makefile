.PHONY: db_apply
db_apply:
	atlas migrate apply -c file://db/atlas.hcl --env local --dir file://db/migrations

.PHONY: gen
gen: gen_server

.PHONY: gen_db
gen_db:
	sqlc generate -f db/sqlc.yaml

.PHONY: gen_server
gen_server:
	mkdir -p "./internal/server"
	oapi-codegen -package server -generate gin,spec openapi.yml > internal/server/server.gen.go
	oapi-codegen -package server -generate types openapi.yml > internal/server/types.gen.go

build:
	go build -o ./app ./cmd/main.go
