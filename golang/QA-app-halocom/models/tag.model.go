package models

// Tag model export
type Tag struct {
	ID uint `json:"id" gorm:"primary_key"`
	TagString string `json:"tag_string"`
}