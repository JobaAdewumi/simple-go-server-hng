package model

import (
	"fmt"
	"html"
	"server/database"
	"strings"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

func (person *Person) Save() (*Person , error) {
	person.Name = html.EscapeString(strings.TrimSpace(person.Name))
	err := database.Database.Create(&person.Name)
    fmt.Print(err)
	if err != nil {
		return &Person{}, err.Error
	}
	// id := err.
	return person, nil
}

func FindPersonById(id string) (Person, error) {
	var person Person
	err := database.Database.Preload("Entries").Where("ID=?", id).Find(&person).Error
	if err != nil {
		return Person{}, err
	}
	return person, nil
}

func DeletePersonById(id string) (string, error) {
    err := database.Database.Delete(&Person{}, id)
    if err != nil {
		return "Delete successful", err.Error
	}
    return "Delete successful", nil
}