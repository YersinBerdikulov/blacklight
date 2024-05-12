package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"Status": "OK",
		"SystemInfo": map[string]string{
			"Environment": app.config.env,
			"Version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
