package controllers

import (
	"fullstack/api/responses"
	"net/http"
)

func (ser Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the Home page")
}
