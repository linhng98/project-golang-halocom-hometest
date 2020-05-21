package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// CreateTopicForm binding json POST request
type createTopicForm struct {
	AccountID uint     `json:"account_id" binding:"required"`
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Tag       []string `json:"tag" binding:"required"`
}

// CreateTopic api (POST /api/topic/create)
func CreateTopic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON createTopicForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new react table
	reactObj := models.React{}
	db.Create(&reactObj)

	// create new topic object
	topicObj := models.Topic{AccountID: inputJSON.AccountID, ReactID: reactObj.ID, Title: inputJSON.Title, Content: inputJSON.Content}
	db.Create(&topicObj)

	for _, s := range inputJSON.Tag {
		// search tag exist or not
		tagObj := models.Tag{TagString: s}
		if err := db.Where("tag_string = ?", s).First(&tagObj).Error; err != nil { // not exist
			// create new tag record
			db.Create(&tagObj)
		}

		// get id from tag row, create new topic-tag record
		topicTagObj := models.TopicTag{TopicID: topicObj.ID, TagID: tagObj.ID}
		db.Create(&topicTagObj)
	}

	c.JSON(http.StatusOK, &topicObj)
}
