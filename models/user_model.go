package models

// User structure
type User struct {
	Username string `json:"username" form:"username" gorm:"primaryKey"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}
