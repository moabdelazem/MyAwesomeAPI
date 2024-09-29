package main

// func (app *Application) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
// 	// Call the storage layer's GetPosts method to get all the posts
// 	posts, err := app.Storage.GetPosts(r.Context())
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}

// 	// Marshal the posts into JSON
// 	err = app.writeJSON(w, http.StatusOK, posts)
// 	if err != nil {
// 		app.errorJSON(w, err)
// 		return
// 	}
// }
