package main

import (
	"fmt"
	"github.com/YersinBerdikulov/blacklight/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Matrix",
		Runtime:   100,
		Genres:    []string{"Science-Fiction", "Thriller", "Action"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Server encountered an internal error and could not process your request", http.StatusInternalServerError)
	}
}
