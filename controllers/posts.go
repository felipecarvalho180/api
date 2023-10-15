package controllers

import (
	"devbook-api/database"
	"devbook-api/helpers"
	"devbook-api/models"
	"devbook-api/repositories"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err := post.Prepare(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, post)
}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPosts(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, posts)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["post_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post, err := repository.GetPostByID(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["post_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postFromDB, err := repository.GetPostByID(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postFromDB.AuthorID != userID {
		helpers.Error(w, http.StatusForbidden, errors.New("you are not allowed to edit this post"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := post.Prepare(); err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.Update(postID, post); err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusNoContent, nil)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := helpers.ExtractUserID(r)
	if err != nil {
		helpers.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["post_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	postFromDB, err := repository.GetPostByID(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postFromDB.AuthorID != userID {
		helpers.Error(w, http.StatusForbidden, errors.New("you are not allowed to delete this post"))
		return
	}

	if err := repository.Delete(postID); err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusNoContent, nil)
}

func GetPostByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.FindByUserID(userID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, posts)
}

func Like(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["post_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	err = repository.Like(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusNoContent, nil)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["post_id"], 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	err = repository.Unlike(postID)
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusNoContent, nil)
}
