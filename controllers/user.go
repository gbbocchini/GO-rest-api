package controllers

import (
	"encoding/json"
	"net/http"
	"rest/database/daos"
	"rest/models"
	"rest/utils"
)

type UserController struct{}

type ResponseOutput struct {
	User  models.User
	Token string
}

func (u UserController) SignupUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		json.NewDecoder(r.Body).Decode(&user)

		if len(user.Name) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Name should be at least 3 characters long!"})
			return
		}

		if len(user.Username) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Username should be at least 3 characters long!"})
			return
		}

		if len(user.Email) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Email should be at least 3 characters long!"})
			return
		}

		if len(user.Password) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Password should be at least 3 characters long!"})
			return
		}

		userDao := daos.UsersDao{}

		_, err := userDao.CreateUser(user)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Failed To Add new user in database!"})
			return
		}

		payload := utils.Payload{
			Username: user.Username,
			Email:    user.Email,
			Id:       user.ID,
		}

		token, err := utils.GenerateJwtToken(payload)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Failed To Generate New JWT Token!"})
			return
		}

		user.Password = ""

		utils.SendSuccess(w, ResponseOutput{
			Token: token,
			User:  user,
		})
	}
}

func (u UserController) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials map[string]string
		json.NewDecoder(r.Body).Decode(&credentials)

		if len(credentials["id"]) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Invalid Username/Email!"})
			return
		}

		if len(credentials["password"]) < 3 {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Invalid Password!"})
			return
		}
		userDao := daos.UsersDao{}
		user, err := userDao.GetUser(credentials)
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Invalid Username/Email, Please Signup!"})
			return
		}

		if user.Password != credentials["password"] {
			utils.SendError(w, http.StatusNotFound, models.Error{Message: "Invalid Credentials!"})
			return
		}

		payload := utils.Payload{
			Username: user.Username,
			Email:    user.Email,
			Id:       user.ID,
		}

		token, err := utils.GenerateJwtToken(payload)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Failed To Generate New JWT Token!"})
			return
		}

		user.Password = ""

		utils.SendSuccess(w, ResponseOutput{
			Token: token,
			User:  user,
		})
	}
}
