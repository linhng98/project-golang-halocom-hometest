package models

// AccountUpvote model export
type AccountUpvote struct {
	ID uint `json:"id" gorm:"primary_key"`
	AccountID uint `json:"account_id"`
	ReactID uint `json:"react_id"`
}