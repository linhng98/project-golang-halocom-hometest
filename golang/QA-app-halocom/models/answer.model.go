package models

// Answer model export
type Answer struct {
	ID uint `json:"id" gorm:"primary_key"`
	TopicID uint `json:"topic_id"`
	AccountID uint `json:"account_id"`
	ReactID uint `json:"react_id" gorm:"unique"` 
	Content string `json:"content" gorm:"type:mediumtext"`
}