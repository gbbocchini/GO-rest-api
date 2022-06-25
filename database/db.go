package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"rest/models"
)

func GetDatabaseAndContext() (*bun.DB, context.Context) {
	var DB *bun.DB
	var DbContext context.Context
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_SQL_URL"))))
	DB = bun.NewDB(sqlDB, pgdialect.New())
	DbContext = context.Background()
	return DB, DbContext
}

func CreateTablesAndPrePopulate() {
	Db, DbContext := GetDatabaseAndContext()
	_, err := Db.NewCreateTable().Model((*models.User)(nil)).Exec(DbContext)
	_, errr := Db.NewCreateTable().Model((*models.Book)(nil)).Exec(DbContext)

	if err != nil && errr != nil {
		fmt.Println("Database Tables already exist. Starting server....")
		return
	}

	books := []models.Book{
		{Title: "Book1", Author: "Author1", Year: 1990},
		{Title: "Book2", Author: "Author2", Year: 2022},
	}

	Db.NewInsert().Model(&books).Exec(DbContext)

}
