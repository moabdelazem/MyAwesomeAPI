package main

import (
	"net/http"

	"github.com/moabdelazem/sample-app/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *Application) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new RegisterUserPayload
	var payload RegisterUserPayload

	// Parse the JSON body into the payload
	if err := ParseJSONBody(w, r, &payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	if err := Validate.Struct(payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create a new User
	user := store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// Hash the password
	if err := user.Password.Set(payload.Password); err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()

	// Call the storage layer's CreateUser method to create the user
	err := app.Storage.Users.CreateUser(ctx, &user)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write the user to the response
	WriteJSON(w, http.StatusCreated, user)
}
