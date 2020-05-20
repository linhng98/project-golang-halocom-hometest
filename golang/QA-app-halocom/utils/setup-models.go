package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/nobabykill/project-golang-halocom-hometest/models"
)

// SetupModels export
func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:halocom@(mariadb-halo)/QA_service?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
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

	db.Model(&models.TopicTag{}).AddUniqueIndex("topic_id", "tag_id")
	db.Model(&models.AccountUpvote{}).AddUniqueIndex("account_id", "react_id")
	db.Model(&models.AccountDownvote{}).AddUniqueIndex("account_id", "react_id")
	db.Model(&models.AccountReport{}).AddUniqueIndex("account_id", "react_id")

	db.Model(&models.Answer{}).AddForeignKey("topic_id", "topics(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Answer{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Answer{}).AddForeignKey("react_id", "reacts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Topic{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Topic{}).AddForeignKey("react_id", "reacts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.TopicTag{}).AddForeignKey("topic_id", "topics(id)", "CASCADE", "RESTRICT")
	db.Model(&models.TopicTag{}).AddForeignKey("tag_id", "tags(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountUpvote{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountUpvote{}).AddForeignKey("react_id", "reacts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountDownvote{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountDownvote{}).AddForeignKey("react_id", "reacts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountReport{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.Model(&models.AccountReport{}).AddForeignKey("react_id", "reacts(id)", "CASCADE", "RESTRICT")

	return db
}
