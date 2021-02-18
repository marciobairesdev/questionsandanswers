package controllers

import "github.com/marciobairesdev/questionsandanswers/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Questions routes
	s.Router.HandleFunc("/questions", middlewares.SetMiddlewareJSON(s.GetQuestions)).Methods("GET")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(s.GetQuestion)).Methods("GET")
	s.Router.HandleFunc("/questions/user/{userId}", middlewares.SetMiddlewareJSON(s.GetQuestionByUser)).Methods("GET")
	s.Router.HandleFunc("/questions", middlewares.SetMiddlewareJSON(s.CreateQuestion)).Methods("POST")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(s.UpdateQuestion)).Methods("PUT")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(s.DeleteQuestion)).Methods("DELETE")
}
