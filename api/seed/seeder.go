package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/marciobairesdev/questionsandanswers/api/models"
)

var questions = []models.Question{
	{
		UserID:    1,
		Statement: "What color is a mirror?",
		Answer:    "",
	},
	{
		UserID:    1,
		Statement: "If anything is possible, is it possible for anything to be impossible?",
		Answer:    "It’s likely to be unlikely, but who really knows?",
	},
	{
		UserID:    2,
		Statement: "Why is a boxing ring square?",
		Answer:    "",
	},
	{
		UserID:    3,
		Statement: "Why the sky is blue?",
		Answer:    "",
	},
}

func Load(db *gorm.DB) {
	err := db.DropTableIfExists(&models.Question{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.AutoMigrate(&models.Question{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}
	for i := range questions {
		err := db.Model(&models.Question{}).Create(&questions[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed questions table: %v", err)
		}
	}
}
