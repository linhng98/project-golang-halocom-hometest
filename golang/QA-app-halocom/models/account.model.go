package models

// Account model export
type Account struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Username       string `json:"username" gorm:"unique"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email" gorm:"unique"`
	Status         bool   `json:"status" gorm:"default true"`
}
