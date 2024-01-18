package models

type Request struct {
	UserID int `json:"userId" validate:"required"`
}

type Response struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	UserRole  string `json:"userRole"`
}
