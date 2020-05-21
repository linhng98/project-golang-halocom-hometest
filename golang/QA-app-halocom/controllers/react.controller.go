package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// ReactForm binding
type ReactForm struct {
	AccountID uint `json:"account_id" binding:"required"`
	ReactID   uint `json:"react_id" binding:"required"`
}

// Upvote api
func Upvote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON ReactForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if this user already upvote for this react or not
	accID := inputJSON.AccountID
	reactID := inputJSON.ReactID
	if err := db.Where("account_id = ? AND react_id = ?", accID, reactID).First(&models.AccountUpvote{}).Error; err != nil {
		// record not found
		// create new record in upvote table
		db.Create(&models.AccountUpvote{
			AccountID: accID,
			ReactID:   reactID})
		// delete in downvote and report
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountReport{})
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountDownvote{})
	} else {
		// record exist
		// delete upvote record
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountUpvote{})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// Downvote api
func Downvote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON ReactForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if this user already downvote for this react or not
	accID := inputJSON.AccountID
	reactID := inputJSON.ReactID
	if err := db.Where("account_id = ? AND react_id = ?", accID, reactID).First(&models.AccountDownvote{}).Error; err != nil {
		// record not found
		// create new record in downvote table
		db.Create(&models.AccountDownvote{
			AccountID: accID,
			ReactID:   reactID})
		// delete in upvote and report
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountReport{})
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountUpvote{})
	} else {
		// record exist
		// delete downvote record
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountDownvote{})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// Report api
func Report(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON ReactForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if this user already report for this react or not
	accID := inputJSON.AccountID
	reactID := inputJSON.ReactID
	if err := db.Where("account_id = ? AND react_id = ?", accID, reactID).First(&models.AccountReport{}).Error; err != nil {
		// record not found
		// create new record in report table
		db.Create(&models.AccountReport{
			AccountID: accID,
			ReactID:   reactID})
		// delete in upvote and downvote
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountUpvote{})
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountDownvote{})
	} else {
		// record exist
		// delete report record
		db.Where("account_id = ? AND react_id = ?", accID, reactID).Delete(&models.AccountReport{})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
