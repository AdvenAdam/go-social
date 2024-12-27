package main

import (
	"net/http"

	"github.com/AdvenAdam/go-social/internal/store"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))

	app.store.Posts.Create(r.Context(), &store.Post{})
}
