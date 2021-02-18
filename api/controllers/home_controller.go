package controllers

import (
	"net/http"

	"github.com/marciobairesdev/questionsandanswers/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.ToJson(w, http.StatusOK, "Welcome to the Question and Answers REST API")
}
