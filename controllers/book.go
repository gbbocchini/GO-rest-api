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

type Controller struct{}

func (c Controller) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		booksDao := daos.BookDAO{}
		books := booksDao.GetBooks()
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.ParseInt(params["id"], 10, 0)

		if err != nil {
			utils.SendError(w, 400, models.Error{Message: "id must be int convertible"})
		}

		booksDao := daos.BookDAO{}
		book := booksDao.GetBookById(id)
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newBookMap map[string]string
		err := json.NewDecoder(r.Body).Decode(&newBookMap)

		if err != nil {
			utils.SendError(w, 400, models.Error{Message: "all values must be strings"})
		}

		newBookYear, _ := strconv.ParseInt(newBookMap["year"], 10, 0)
		newBook := models.Book{
			Title:  newBookMap["title"],
			Author: newBookMap["author"],
			Year:   newBookYear,
		}
		booksDao := daos.BookDAO{}
		newlyAdded := booksDao.AddBook(newBook)

		utils.SendSuccess(w, newlyAdded)
	}
}

func (c Controller) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var updateData map[string]string
		err := json.NewDecoder(r.Body).Decode(&updateData)

		if err != nil {
			utils.SendError(w, 400, models.Error{Message: "all values must be strings"})
		}

		bookId, errr := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)
		bookToUpdate := models.Book{ID: bookId}

		if errr != nil {
			utils.SendError(w, 400, models.Error{Message: "impossible to parse id"})
		}

		if updateData["title"] != "" {
			bookToUpdate.Title = updateData["title"]
		}

		if updateData["author"] != "" {
			bookToUpdate.Author = updateData["author"]
		}

		if updateData["year"] != "" {
			newBookYear, err := strconv.ParseInt(updateData["year"], 10, 0)
			if err != nil {
				utils.SendError(w, 400, models.Error{Message: "impossible to parse year"})
			}
			bookToUpdate.Year = newBookYear
		}

		booksDao := daos.BookDAO{}
		newlyAdded := booksDao.UpdateBook(bookToUpdate)
		utils.SendSuccess(w, newlyAdded)
	}
}

func (c Controller) RemoveBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToRemove, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)

		if err != nil {
			utils.SendError(w, 400, models.Error{Message: "impossible to parse id"})
		}

		booksDao := daos.BookDAO{}
		deleted := booksDao.DeleteBook(idToRemove)
		utils.SendSuccess(w, deleted)
	}
}
