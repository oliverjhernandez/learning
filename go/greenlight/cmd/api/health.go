package main

import (
	"net/http"
)

func (a *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	version := "0.0.1"
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.env,
			"version":     version,
		},
	}

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.internalServerError(w, r)
	}
}
