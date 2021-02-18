package models

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/marciobairesdev/questionsandanswers/api/controllers"
	"github.com/marciobairesdev/questionsandanswers/api/models"
	"github.com/marciobairesdev/questionsandanswers/api/seed"
	"gopkg.in/go-playground/assert.v1"
)

var server = controllers.Server{}
var questionInstance = models.Question{}

func init() {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_NAME"))
	seed.Load(server.DB)
}

func TestFindAllQuestions(t *testing.T) {
	questions, err := questionInstance.FindAllQuestions(server.DB)
	if err != nil {
		t.Errorf("This is the error finding questions: %v\n", err)
		return
	}
	assert.Equal(t, len(*questions), 4)
}

func TestFindQuestionByID(t *testing.T) {
	foundQuestion, err := questionInstance.FindQuestionByID(server.DB, 2)
	if err != nil {
		t.Errorf("This is the error finding question by ID: %v\n", err)
		return
	}
	assert.Equal(t, foundQuestion.ID, 2)
	assert.Equal(t, foundQuestion.UserID, 1)
	assert.Equal(t, foundQuestion.Statement, "If anything is possible, is it possible for anything to be impossible?")
	assert.Equal(t, foundQuestion.Answer, "It’s likely to be unlikely, but who really knows?")
}

func TestFindQuestionsByUserID(t *testing.T) {
	userQuestions, err := questionInstance.FindQuestionsByUserID(server.DB, 1)
	if err != nil {
		t.Errorf("This is the error finding questions by user ID: %v\n", err)
		return
	}
	assert.Equal(t, len(*userQuestions), 2)
}

func TestCreateQuestion(t *testing.T) {
	newQuestion := models.Question{
		UserID:    3,
		Statement: "How big is the universe?",
		Answer:    "Only GOD knows...",
	}
	createdQuestion, err := newQuestion.CreateQuestion(server.DB)
	if err != nil {
		t.Errorf("This is the error creating a new question: %v\n", err)
		return
	}
	assert.Equal(t, createdQuestion.ID, 5)
	assert.Equal(t, createdQuestion.UserID, newQuestion.UserID)
	assert.Equal(t, createdQuestion.Statement, newQuestion.Statement)
	assert.Equal(t, createdQuestion.Answer, newQuestion.Answer)
}

func TestUpdateQuestion(t *testing.T) {
	foundQuestion, err := questionInstance.FindQuestionByID(server.DB, 2)
	if err != nil {
		t.Errorf("This is the error finding question 2 for update: %v\n", err)
		return
	}
	newQuestion := models.Question{
		ID:        foundQuestion.ID,
		UserID:    foundQuestion.UserID,
		Statement: foundQuestion.Statement,
		Answer:    "The color of the soul of those who look at it!",
		CreatedAt: foundQuestion.CreatedAt,
		UpdatedAt: foundQuestion.UpdatedAt,
	}
	createdQuestion, err := newQuestion.UpdateQuestion(server.DB, uint32(newQuestion.ID))
	if err != nil {
		t.Errorf("This is the error updating question 2: %v\n", err)
		return
	}
	assert.Equal(t, createdQuestion.ID, foundQuestion.ID)
	assert.Equal(t, createdQuestion.UserID, foundQuestion.UserID)
	assert.Equal(t, createdQuestion.Statement, foundQuestion.Statement)
	assert.Equal(t, createdQuestion.Answer, newQuestion.Answer)
	assert.Equal(t, createdQuestion.CreatedAt, foundQuestion.CreatedAt)
	assert.NotEqual(t, createdQuestion.UpdatedAt, foundQuestion.UpdatedAt)
}

func TestDeleteQuestion(t *testing.T) {
	isDeleted, err := questionInstance.DeleteQuestion(server.DB, 4)
	if err != nil {
		t.Errorf("This is the error deleting question 4: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}
