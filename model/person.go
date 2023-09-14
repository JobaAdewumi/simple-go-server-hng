package model

import "gorm.io/gorm"

type Person struct {
    gorm.Model
    Name string `gorm:"size:255;not null;unique" json:"name"`
}