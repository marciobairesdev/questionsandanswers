package models

import (
	"errors"
	"fmt"
	"html"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Question struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	UserID    int       `gorm:"index" json:"user_id" validate:"required"`
	Statement string    `gorm:"size:255;not null;" json:"statement" validate:"required"`
	Answer    string    `gorm:"size:255" json:"answer,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" validate:"required"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" validate:"required"`
}

func (q *Question) validate() error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	errs := validate.Struct(*q)
	if errs == nil {
		return nil
	}
	for _, err := range errs.(validator.ValidationErrors) {
		return err
	}
	return nil
}

// FindAllQuestions ...
func (q *Question) FindAllQuestions(db *gorm.DB) (*[]Question, error) {
	var err error
	questions := []Question{}
	err = db.Model(&Question{}).Find(&questions).Error
	if err != nil {
		return &[]Question{}, err
	}
	return &questions, err
}

// FindQuestionByID ...
func (q *Question) FindQuestionByID(db *gorm.DB, uid uint32) (*Question, error) {
	var err error
	err = db.Model(Question{}).Where("id = ?", uid).Take(&q).Error
	if err != nil {
		return &Question{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Question{}, errors.New("Question not found")
	}
	return q, err
}

// FindQuestionByUserID ...
func (q *Question) FindQuestionsByUserID(db *gorm.DB, userUID uint32) (*[]Question, error) {
	var err error
	userQuestions := []Question{}
	err = db.Model(Question{}).Where("user_id = ?", userUID).Find(&userQuestions).Error
	if err != nil {
		return &[]Question{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Question{}, errors.New(fmt.Sprintf("Question(s) for user %d not found", userUID))
	}
	return &userQuestions, err
}

// CreateQuestion ...
func (q *Question) CreateQuestion(db *gorm.DB) (*Question, error) {
	q.ID = 0
	q.Statement = html.EscapeString(strings.TrimSpace(q.Statement))
	q.Answer = html.EscapeString(strings.TrimSpace(q.Answer))
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	if err := q.validate(); err != nil {
		return &Question{}, err
	}
	var err error
	err = db.Create(&q).Error
	if err != nil {
		return &Question{}, err
	}
	return q, nil
}

// UpdateQuestion ...
func (q *Question) UpdateQuestion(db *gorm.DB, uid uint32) (*Question, error) {
	if err := q.validate(); err != nil {
		return &Question{}, err
	}
	db = db.Model(&Question{}).Where("id = ?", uid).Take(&Question{}).UpdateColumns(
		map[string]interface{}{
			"statement":  html.EscapeString(strings.TrimSpace(q.Statement)),
			"answer":     html.EscapeString(strings.TrimSpace(q.Answer)),
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Question{}, db.Error
	}
	err := db.Model(&Question{}).Where("id = ?", uid).Take(&q).Error
	if err != nil {
		return &Question{}, err
	}
	return q, nil
}

// DeleteQuestion ...
func (q *Question) DeleteQuestion(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Model(&Question{}).Where("id = ?", uid).Take(&Question{}).Delete(&Question{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
