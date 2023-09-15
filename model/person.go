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
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}



func (person *Person) Save() (*Person, error) {
	person.Name = html.EscapeString(strings.TrimSpace(person.Name))
	err := database.Database.Select("Name").Create(&person)
	fmt.Print(err)
	fmt.Print(person)

	if err.Error != nil {
		return &Person{}, err.Error
	}
	// id := err.
	return person, nil
}

func (p *Person) Update() (int64, error) {
	var person Person
	// person.Name = html.EscapeString(strings.TrimSpace(person.Name))
	err := database.Database.Save(&p)
	fmt.Print(err)
	fmt.Print(person)
	fmt.Print(p)
	if err.Error != nil {
		return 0, err.Error
	}
	// id := err.
	return err.RowsAffected, nil
}

func FindPersonById(id string) (Person, error) {
	var person Person
	err := database.Database.Find(&person, id)
	if err.Error != nil {
		return Person{}, err.Error
	}
	return person, nil
}

func UpdatePersonById(id string) error {
	var person Person
	err := database.Database.Find(&person, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeletePersonById(id string) (string, error) {
	// var person Person
	err := database.Database.Unscoped().Delete(&Person{}, id)
	fmt.Print(err)
	if err.RowsAffected == 0 {
		return "Person cannot be found", err.Error
	}
	if err.Error != nil {
		return "Error Deleting", err.Error
	}
	return "Delete successful", nil
}
