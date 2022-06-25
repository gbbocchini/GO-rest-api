package daos

import (
	"rest/database"
	"rest/models"
)

type BookDAO struct{}

const DatabaseError = "Database error. Please contact support."

func (b BookDAO) GetBooks() ([]models.Book, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	var books []models.Book
	err := Db.NewSelect().Model((*models.Book)(nil)).Scan(DbContext, &books)

	return books, err
}

func (b BookDAO) GetBookById(id int64) (models.Book, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	var book models.Book
	err := Db.NewSelect().Model(&book).Where("id = ?", id).Scan(DbContext)
	return book, err
}

func (b BookDAO) AddBook(data models.Book) (models.Book, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	_, err := Db.NewInsert().Model(&data).Exec(DbContext)
	return data, err
}

func (b BookDAO) UpdateBook(updateBook models.Book) (models.Book, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	_, err := Db.NewUpdate().Model(&updateBook).OmitZero().WherePK().Exec(DbContext)
	var updatedBook models.Book
	Db.NewSelect().Model(&updatedBook).Where("id = ?", updatedBook.ID).Scan(DbContext)
	return updatedBook, err
}

func (b BookDAO) DeleteBook(id int64) (int64, error) {
	Db, DbContext := database.GetDatabaseAndContext()
	toDelete := models.Book{ID: id}
	_, err := Db.NewDelete().Model(&toDelete).WherePK().Exec(DbContext)
	return id, err
}
