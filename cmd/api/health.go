package main

import (
	"net/http"
)

// healthCheck godoc
//
//	@Summary		Checks the health of the API
//	@Description	Checks the health of the API
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/users/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": app.config.version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
