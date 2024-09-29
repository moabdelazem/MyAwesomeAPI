package main

import "net/http"

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
