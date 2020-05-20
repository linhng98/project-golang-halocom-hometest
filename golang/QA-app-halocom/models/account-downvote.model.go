package models

// AccountDownvote model export
type AccountDownvote struct {
	ID uint `json:"id" gorm:"primary_key"`
	AccountID uint `json:"account_id"`
	ReactID uint `json:"react_id"`
}