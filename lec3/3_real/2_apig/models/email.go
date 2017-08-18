package models

type Email struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	UserID  uint   `json:"user_id" form:"user_id"`
	Address string `json:"address" form:"address"`
	User    *User  `json:"user form:"user`
}
