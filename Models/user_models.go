package models

// User structure
type User struct {
	Picture  []byte `json:"pic"`
	Fullname string `json:"fullname"`
	Contact  string `json:"contact"`
}
