package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/marciobairesdev/questionsandanswers/api/models"
	"github.com/marciobairesdev/questionsandanswers/api/responses"
)

func (server *Server) GetQuestions(w http.ResponseWriter, r *http.Request) {
	question := models.Question{}
	questions, err := question.FindAllQuestions(server.DB)
	if err != nil {
		responses.ToError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ToJson(w, http.StatusOK, questions)
}

func (server *Server) GetQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	question := models.Question{}
	questionGotten, err := question.FindQuestionByID(server.DB, uint32(uid))
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	responses.ToJson(w, http.StatusOK, questionGotten)
}

func (server *Server) GetQuestionByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userUID, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	question := models.Question{}
	userQuestions, err := question.FindQuestionsByUserID(server.DB, uint32(userUID))
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	responses.ToJson(w, http.StatusOK, userQuestions)
}

func (server *Server) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
	}
	question := models.Question{}
	err = json.Unmarshal(body, &question)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	questionCreated, err := question.CreateQuestion(server.DB)
	if err != nil {
		responses.ToError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, questionCreated.ID))
	responses.ToJson(w, http.StatusCreated, questionCreated)
}

func (server *Server) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	question := models.Question{}
	err = json.Unmarshal(body, &question)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	if question.ID != int(uid) {
		responses.ToError(w, http.StatusInternalServerError, errors.New("Path ID mismatch from payload ID"))
		return
	}
	questionGotten := &models.Question{}
	questionGotten, err = questionGotten.FindQuestionByID(server.DB, uint32(uid))
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	question.UserID = questionGotten.UserID
	question.CreatedAt = questionGotten.CreatedAt
	question.UpdatedAt = questionGotten.UpdatedAt
	if len(strings.TrimSpace(question.Statement)) == 0 && len(strings.TrimSpace(questionGotten.Statement)) > 0 {
		question.Statement = questionGotten.Statement
	}
	updatedQuestion, err := question.UpdateQuestion(server.DB, uint32(uid))
	if err != nil {
		responses.ToError(w, http.StatusInternalServerError, err)
		return
	}
	responses.ToJson(w, http.StatusOK, updatedQuestion)
}

func (server *Server) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	question := models.Question{}
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ToError(w, http.StatusBadRequest, err)
		return
	}
	_, err = question.DeleteQuestion(server.DB, uint32(uid))
	if err != nil {
		responses.ToError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.ToJson(w, http.StatusNoContent, "")
}
