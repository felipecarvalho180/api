package controllers

import (
	"devbook-api/database"
	"devbook-api/helpers"
	"devbook-api/models"
	"devbook-api/repositories"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userDB, err := repository.FindByEmail(user.Email)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = helpers.VerifyPassword(user.Password, userDB.Password); err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := helpers.GenerateToken(userDB.ID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, token)
}
