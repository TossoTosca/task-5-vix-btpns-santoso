package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakamin-final-project/core"
	"github.com/rakamin-final-project/helpers"
)

type UserController struct {
	UserService core.UserService
}

func NewUserController(userService core.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = uc.UserService.CreateUser(&user)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusCreated, user)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginData core.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := uc.UserService.Login(&loginData)
	if err != nil {
		helpers.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := uc.UserService.GetUser(userID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = uc.UserService.UpdateUser(userID, &user)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, user)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	err := uc.UserService.DeleteUser(userID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
