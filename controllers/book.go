package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest/database/daos"
	"rest/models"
	"rest/utils"
	"strconv"
)

type BookController struct{}

func (c BookController) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		booksDao := daos.BookDAO{}
		books, err := booksDao.GetBooks()

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: daos.DatabaseError})
			return
		}

		utils.SendSuccess(w, books)
	}
}

func (c BookController) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["id"], 10, 0)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "id must be int convertible"})
			return
		}

		booksDao := daos.BookDAO{}
		book, err := booksDao.GetBookById(id)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: daos.DatabaseError})
			return
		}

		utils.SendSuccess(w, book)
	}
}

func (c BookController) AddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBookMap map[string]string
		err := json.NewDecoder(r.Body).Decode(&newBookMap)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "all values must be strings"})
			return
		}

		if newBookMap["title"] == "" || len(newBookMap["title"]) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "title is mandatory"})
			return
		}

		if newBookMap["author"] == "" || len(newBookMap["author"]) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "author is mandatory"})
			return
		}

		if newBookMap["utils"] == "" || len(newBookMap["year"]) < 4 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "year is mandatory"})
			return
		}

		newBookYear, _ := strconv.ParseInt(newBookMap["year"], 10, 0)
		newBook := models.Book{
			Title:  newBookMap["title"],
			Author: newBookMap["author"],
			Year:   newBookYear,
		}
		booksDao := daos.BookDAO{}
		newlyAdded, err := booksDao.AddBook(newBook)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: daos.DatabaseError})
			return
		}

		utils.SendSuccess(w, newlyAdded)
	}
}

func (c BookController) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateData map[string]string
		err := json.NewDecoder(r.Body).Decode(&updateData)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "all values must be strings"})
			return
		}

		bookId, errr := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)

		if errr != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "impossible to parse id"})
			return
		}

		bookToUpdate := models.Book{ID: bookId}

		if len(updateData["title"]) > 3 {
			bookToUpdate.Title = updateData["title"]
		} else {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "title minimum length is 4"})
			return
		}

		if len(updateData["author"]) > 3 {
			bookToUpdate.Author = updateData["author"]
		} else {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "author minimum length is 4"})
			return
		}

		if len(updateData["year"]) > 4 {
			newBookYear, err := strconv.ParseInt(updateData["year"], 10, 0)

			if err != nil {
				utils.SendError(w, http.StatusBadRequest, models.Error{Message: "impossible to parse year"})
				return
			}

			bookToUpdate.Year = newBookYear
		} else {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "year minimum length is 4"})
			return
		}

		booksDao := daos.BookDAO{}
		newlyAdded, errrr := booksDao.UpdateBook(bookToUpdate)

		if errrr != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: daos.DatabaseError})
			return
		}

		utils.SendSuccess(w, newlyAdded)
	}
}

func (c BookController) RemoveBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToRemove, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "impossible to parse id"})
			return
		}

		booksDao := daos.BookDAO{}
		deleted, errr := booksDao.DeleteBook(idToRemove)

		if errr != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: daos.DatabaseError})
			return
		}

		utils.SendSuccess(w, deleted)
	}
}
