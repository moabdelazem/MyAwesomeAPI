package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetUsersHandler is a method on the Application struct that handles the GET /v1/users route
//   - This method is responsible for getting all the users from the storage layer and writing them to the response
func (app *Application) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Call the storage layer's GetPosts method to get all the posts
	posts, err := app.Storage.Users.GetUsers(r.Context())
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Marshal the posts into JSON
	err = WriteJSON(w, http.StatusOK, posts)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (app *Application) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get The Target User ID
	id := chi.URLParam(r, "id")

	// Parse the ID into a UUID
	userID, err := uuid.Parse(id)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Call the storage layer's GetUserByID method to get the user
	user, err := app.Storage.Users.GetUserByID(r.Context(), userID)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (app *Application) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Get The Target User ID
	username := chi.URLParam(r, "username")

	// Call the storage layer's GetUserByID method to get the user
	user, err := app.Storage.Users.GetUserByUsername(r.Context(), username)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, user)
}
