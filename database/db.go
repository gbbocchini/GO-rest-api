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
		{Title: "Book1", Author: "Author1", Year: 1983},
		{Title: "Book2", Author: "Author2", Year: 1990},
		{Title: "Book3", Author: "Author3", Year: 2000},
		{Title: "Book4", Author: "Author4", Year: 2010},
		{Title: "Book5", Author: "Author5", Year: 2020},
		{Title: "Book6", Author: "Author6", Year: 2022},
	}

	_, err = Db.NewInsert().Model(&books).Exec(DbContext)

	if err != nil {
		fmt.Println("Could not create sample books on database table 'books', please check...")
		return
	}

}
