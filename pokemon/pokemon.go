package pokemon

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Attack  int    `json:"attack"`
	HP      int    `json:"hp"`
	Defense int    `json:"defense"`
}
