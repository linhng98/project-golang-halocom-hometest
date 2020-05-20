package models

// TopicTag model export
type TopicTag struct {
	ID uint `json:"id" gorm:"primary_key"`
	TopicID uint `json:"topic_id"`
	TagID uint `json:"tag_id"`
}