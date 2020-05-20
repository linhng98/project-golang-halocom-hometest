package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nobabykill/project-golang-halocom-hometest/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	db := utils.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Run()
}
