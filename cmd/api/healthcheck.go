package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status":      "OK",
		"Environment": app.config.env,
		"Version":     version,
	}

	jsonResponse, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered an internal error, could not process your request", http.StatusInternalServerError)
		return
	}

	jsonResponse = append(jsonResponse, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
