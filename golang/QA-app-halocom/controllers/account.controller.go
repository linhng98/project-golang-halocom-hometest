package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// CreateAccountForm binding json POST request
type createAccountForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// CreateAccount api (POST /api/account/create)
func CreateAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate form post json
	var inputJSON createAccountForm
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h := sha256.New()
	h.Write([]byte(inputJSON.Password))                             // create new hash
	hashedBase64PW := base64.StdEncoding.EncodeToString(h.Sum(nil)) // encode base64
	accountObj := models.Account{Username: inputJSON.Username, HashedPassword: hashedBase64PW, Email: inputJSON.Email}

	db.Create(&accountObj)
	c.JSON(http.StatusOK, &accountObj)
}
