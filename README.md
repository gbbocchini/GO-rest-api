# GOLang REST API study implementation

A simple rest-api made with GO, Mux and Bun with authentication using go-jwt.

## Setting up:

First create a new POSTGRESQL database instance. After that create a ```.env``` file with the following env vars:

```
POSTGRES_SQL_URL="postgres://<YOUR_POSTGRES_USER>:<YOUR_POSTGRES_PASSWORD>@<YOUR_POSTGRES_HOST>:<YOUR_POSTGRES_PORT>/<YOUR_DB_INTANCE_NAME>?sslmode=disable"
JWT_SECRET="<YOUR_SUPER_SECRET>"
```
To run locally (since the app uses Go Modules, you can run the command bellow and Go will take care of downloading dependencies).
At first run, the app will check the database and create the necessary tables. In the table books some example books will be inserted.
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
- https://pkg.go.dev/github.com/golang-jwt/jwt
- https://github.com/subosito/gotenv