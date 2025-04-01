package entity

import "time"

type Area struct {
	Cords      []GeoCords
	CreateDate time.Time
	Id         int
	IsActive   bool
	Name       string
	UpdateDate time.Time
}
