package models

// Topic model export
type Topic struct {
	ID uint `json:"id" gorm:"primary_key"`
	AccountID uint `json:"account_id"`
	ReactID uint `json:"react_id" gorm:"unique"`
	Title string `json:"title" gorm:"type:varchar(300)"`
	Content string `json:"content" gorm:"type:mediumtext"` 
}