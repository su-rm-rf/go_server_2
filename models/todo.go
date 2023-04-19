package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id int `json:id`
	Text string	`json:text`
	Completed int `json:completed`
}