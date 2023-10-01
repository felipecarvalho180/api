package controllers

import (
	"devbook-api/database"
	"devbook-api/models"
	"devbook-api/repositories"
	utils "devbook-api/utils/entities"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.SIGN_UP_STEP); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Find(nameOrNick)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, err := repository.FindByID(userID)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.UPDATE_USER_STEP); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Update(userID, user); err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err := repository.Delete(userID); err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}
