# GOLang REST API study implementation

A simple rest-api made with GO, Mux and Bun
## Setting up for Development

First create a POSTGRESQL instance and inside it, a table called books (columns: id bigserial, title varchar, 
author varchar, year int8). Then export the db URI:

```
export POSTGRES_SQL_URL="postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@<POSTGRES_HOST>:<POSTGRES_PORT>/<DB_INTANCE_NAME>?sslmode=disable"
```
To run locally:
```
go run main.go
```
To build:
```
go build
```
To run the build:
```
./rest
```

## References
- https://github.com/gorilla/mux
- https://bun.uptrace.dev/