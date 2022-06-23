package daos

import (
	"context"
	"github.com/uptrace/bun"
	"rest/database"
	"rest/models"
)

type BookDAO struct{}

var db *bun.DB
var dbContext context.Context

func (b BookDAO) GetBooks() []models.Book {
	db, dbContext = database.GetDatabaseAndContext()
	var books []models.Book
	db.NewSelect().Model((*models.Book)(nil)).Scan(dbContext, &books)
	return books
}

func (b BookDAO) GetBookById(id int64) models.Book {
	db, dbContext = database.GetDatabaseAndContext()
	var book models.Book
	db.NewSelect().Model(&book).Where("id = ?", id).Scan(dbContext)
	return book
}

func (b BookDAO) AddBook(data models.Book) models.Book {
	db, dbContext = database.GetDatabaseAndContext()
	db.NewInsert().Model(&data).Exec(dbContext)
	return data
}

func (b BookDAO) UpdateBook(updateBook models.Book) models.Book {
	db, dbContext = database.GetDatabaseAndContext()
	db.NewUpdate().Model(&updateBook).OmitZero().WherePK().Exec(dbContext)
	var updatedBook models.Book
	db.NewSelect().Model(&updatedBook).Where("id = ?", updatedBook.ID).Scan(dbContext)
	return updatedBook
}

func (b BookDAO) DeleteBook(id int64) int64 {
	db, dbContext = database.GetDatabaseAndContext()
	toDelete := models.Book{ID: id}
	db.NewDelete().Model(&toDelete).WherePK().Exec(dbContext)
	return id
}
