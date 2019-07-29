package models

// Admin model
type Admin struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}
