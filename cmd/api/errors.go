package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %v", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %v", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found response: %s path: %s error: %v", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflit error: %s path: %s error: %v", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusConflict, err.Error())
}
