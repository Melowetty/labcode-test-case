package entity

import "time"

type Area struct {
	Cords       []GeoCords
	CreatedDate time.Time
	Id          int
	IsActive    bool
	Name        string
	UpdatedDate time.Time
}
