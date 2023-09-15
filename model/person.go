package model

import (
	"fmt"
	"server/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `gorm:"size:255;not null;unique" json:"name"`
}


func (person *Person) Save() (*Person, error) {
	newPerson := Person{
    ID:   uuid.New(),
    Name: person.Name,
}
	err := database.Database.Create(&newPerson)

	if err.Error != nil {
		return &Person{}, err.Error
	}
	return &newPerson, nil
}

func (p *Person) Update() (int64, error) {
	
	var person Person
	// person.Name = html.EscapeString(strings.TrimSpace(person.Name))
	err := database.Database.Session(&gorm.Session{SkipHooks: true}).Save(&p)
	fmt.Print(err)
	fmt.Print(person)
	fmt.Print(p)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func FindPersonById(id string) (Person, error) {
	var person Person
	err := database.Database.Find(&person, "id = ?", id)
	if err.Error != nil {
		return Person{}, err.Error
	}
	return person, nil
}

func UpdatePersonById(id string) error {
	var person Person
	err := database.Database.Find(&person, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeletePersonById(id string) (string, error) {
	err := database.Database.Unscoped().Delete(&Person{}, "id = ?", id)
	if err.RowsAffected == 0 {
		return "Person cannot be found", err.Error
	}
	if err.Error != nil {
		return "Error Deleting", err.Error
	}
	return "Delete successful", nil
}
