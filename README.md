# GOLang REST API study implementation

A simple rest-api made with GO, Mux and Bun with authentication using go-jwt.

## Setting up:

First create a POSTGRESQL instance and inside it. Then, create a .env file with the following env vars:

```
POSTGRES_SQL_URL="postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@<POSTGRES_HOST>:<POSTGRES_PORT>/<DB_INTANCE_NAME>?sslmode=disable"
JWT_SECRET="<yoursupersecret>"
```
To run locally (since the app uses Go Modules, you can run the command bellow and Go will take care of downloading dependencies).
At first run, the app will check the database and create the necessary tables. In the table books, 2 example books will be inserted.
```
go run main.go
```
To build:
```
go build
```
To run the build:
```
./GO-rest-api
```

## References
- https://github.com/gorilla/mux
- https://bun.uptrace.dev/
- https://pkg.go.dev/github.com/golang-jwt/jwt
- https://github.com/subosito/gotenv