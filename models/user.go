package models

//User is a user
type User struct {
	ID       uint   `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Username string `json:"username" gorm:"unique_index"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}
