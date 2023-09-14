package model

type NameInput struct {
	Name string `json:"name" binding:"required"`
}
