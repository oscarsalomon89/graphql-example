package models

import "time"

type Employee struct {
	ID          int
	Name        string
	City        string
	IdRol       int
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}
