package models

// React model export
type React struct {
	ID uint `json:"id" gorm:"primary_key"`
	Upvote uint `json:"upvote" gorm:"default 0"`
	Downvote uint `json:"downvote" gorm:"default 0"`
	Report uint `json:"report" gorm:"default 0"`
}