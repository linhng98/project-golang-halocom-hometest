package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// createAnswerForm binding struct POST request
type createAnswerForm struct {
	TopicID   uint   `json:"topic_id" binding:"required"`
	AccountID uint   `json:"account_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

// CreateAnswer api (POST /api/answer/create)
func CreateAnswer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON createAnswerForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new react table
	reactObj := models.React{}
	db.Create(&reactObj)

	answerObj := models.Answer{
		TopicID:   inputJSON.TopicID,
		AccountID: inputJSON.AccountID,
		ReactID:   reactObj.ID,
		Content:   inputJSON.Content}
	db.Create(&answerObj)

	c.JSON(http.StatusOK, &answerObj)
}

// GetAllAnswer api (POST /api/answer/get-all)
func GetAllAnswer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	topicID := c.Query("topic_id")
	if topicID == "" { // empty string, invalid request
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid topic_id"})
	}

	var answers []models.Answer
	db.Where("topic_id = ?", topicID).Find(&answers)
	c.JSON(http.StatusOK, &answers)
}
