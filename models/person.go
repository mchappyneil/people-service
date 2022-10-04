package models

type Person struct {
	Name    string `json:"name" validate:"required,min=2,max=32"`
	Address string `json:"address" validate:"required"`
}
