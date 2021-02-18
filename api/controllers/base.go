package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/marciobairesdev/questionsandanswers/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	server.DB, err = gorm.Open(dbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbDriver)
		log.Fatal("This is the error:", err)
	}

	server.DB.AutoMigrate(&models.Question{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
