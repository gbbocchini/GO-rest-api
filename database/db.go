package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func GetDatabaseAndContext() (*bun.DB, context.Context) {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_SQL_URL"))))
	DB := bun.NewDB(sqlDB, pgdialect.New())
	DBContext := context.Background()

	return DB, DBContext
}
