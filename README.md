# go-api-demo

This project is to showcase all the common technology I've used over the years building Go services.
This project serves to be the potential backend for my main website hosted at 

# Server Code Generation

Server code is generated from the openapi.yml spec via https://github.com/deepmap/oapi-codegen. When changing the spec run

```commandline
make gen_server
```