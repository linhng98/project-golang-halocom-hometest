package utils

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// SetupModels export
func SetupModels() *gorm.DB {
	var db *gorm.DB
	var err error

	for { // check if database is ready or not
		db, err = gorm.Open("mysql", "root:halocom@(mariadb-halo)/QA_service?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			fmt.Println("Failed to connect to database!")
			time.Sleep(2 * time.Second) // database not ready, sleep 2 second and reconnect
		} else {
			break
		}
	}

	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.React{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Answer{})
	db.AutoMigrate(&models.Topic{})
	db.AutoMigrate(&models.TopicTag{})
	db.AutoMigrate(&models.AccountUpvote{})
	db.AutoMigrate(&models.AccountDownvote{})
	db.AutoMigrate(&models.AccountReport{})

	db.Model(&models.TopicTag{}).AddUniqueIndex("unique_topic_tag_TopicTag", "topic_id", "tag_id")
	db.Model(&models.AccountUpvote{}).AddUniqueIndex("unique_account_react_AccountUpvote", "account_id", "react_id")
	db.Model(&models.AccountDownvote{}).AddUniqueIndex("unique_account_react_AccountDownvote", "account_id", "react_id")
	db.Model(&models.AccountReport{}).AddUniqueIndex("unique_account_react_AccountReport", "account_id", "react_id")

	db.Model(&models.Answer{}).AddForeignKey("topic_id", "topics(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Answer{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Answer{}).AddForeignKey("react_id", "reacts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Topic{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Topic{}).AddForeignKey("react_id", "reacts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.TopicTag{}).AddForeignKey("topic_id", "topics(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.TopicTag{}).AddForeignKey("tag_id", "tags(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountUpvote{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountUpvote{}).AddForeignKey("react_id", "reacts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountDownvote{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountDownvote{}).AddForeignKey("react_id", "reacts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountReport{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.AccountReport{}).AddForeignKey("react_id", "reacts(id)", "RESTRICT", "RESTRICT")

	return db
}
