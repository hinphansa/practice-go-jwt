package models

// Profile structure
type Profile struct {
	Picture  []byte `json:"pic"`
	Fullname string `json:"fullname"`
	Contact  string `json:"contact"`
}
