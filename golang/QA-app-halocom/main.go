package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nobabykill/project-golang-halocom-hometest/controllers"
	"github.com/nobabykill/project-golang-halocom-hometest/utils"
)

func main() {
	r := gin.Default()
	db := utils.SetupModels()

	// Provide db variable to controller
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/api/topic/get-all", controllers.GetAllTopic)
	r.POST("/api/topic/create", controllers.CreateTopic)

	r.POST("/api/account/create", controllers.CreateAccount)

	r.GET("/api/answer/get-all", controllers.GetAllAnswer)
	r.POST("/api/answer/create", controllers.CreateAnswer)

	r.POST("/api/react/upvote", controllers.Upvote)
	r.POST("/api/react/downvote", controllers.Downvote)
	r.POST("/api/react/report", controllers.Report)

	r.Run()
}
