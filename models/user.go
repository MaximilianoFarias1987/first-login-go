package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	LastName    string    `json:"lastName"`
	Address     string    `json:"adrress"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	UserName    string    `json:"userName"`
	Password    string    `json:"password"`
}
