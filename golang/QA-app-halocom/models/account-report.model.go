package models

// AccountReport model export
type AccountReport struct {
	ID        uint `json:"id" gorm:"primary_key"`
	AccountID uint `json:"account_id"`
	ReactID   uint `json:"react_id"`
}
