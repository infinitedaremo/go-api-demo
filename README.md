# go-api-demo

This project is to showcase all the common technology I've used over the years building Go services.
This project serves to be the potential backend for my main website hosted at 

## Dependencies (brew)
* sqlc
* sqlite
* [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
* [mockery](https://vektra.github.io/mockery)

## Server Code Generation

Server code is generated from the openapi.yml spec via https://github.com/deepmap/oapi-codegen. 
When changing the spec run the following

```commandline
make gen_server
```

This follows API-First Development practices and reduces the boilerplate needed for setting up an API server. Output 
for the auto generated server code is at [internal/server](internal/server)

## Database Layer Code Generation

Database interaction is generated by [sqlc](https://sqlc.dev/). All that is defined is the SQL queries with sqlc
definitions for interaction and the app code for interaction is auto generated and safety checked. This allows for frequent
schema changes over time ensuring no future changes would impact production services.

For simplicity’s sake and for demoing, this project utilises sqllite which to run locally requires github.com/mattn/go-sqlite3
installed on your machine. sqlc also needs to be installed on your machine

The auto generated code is at [internal/db](internal/db)

## Services

The service layer is split into a Controller-Service due to the size of the project. For a larger scale project
I would adopt the traditional Service/Handler pattern, splitting up business and service layer logic for ease of testing.

## Hosting

The demo is hosted at https://go-api-demo.onrender.com for testing. If inactive for some time it will take about
50 seconds to re-deploy on free tier.

Samples:

```commandline
curl https://go-api-demo.onrender.com/ping
curl https://go-api-demo.onrender.com/portfolio/1
```

## Mocking

This app leverages [mockery](https://vektra.github.io/mockery) for ease of mock generations for interfaces 