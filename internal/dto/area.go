package dto

import "time"

type AreaDetailed struct {
	Cords       []GeoCords `json:"cords"`
	CreatedDate time.Time  `json:"created_date"`
	Id          int        `json:"id"`
	IsActive    bool       `json:"is_active"`
	Name        string     `json:"name"`
	UpdatedDate time.Time  `json:"updated_date"`
	Cameras     []Camera   `json:"cameras"`
}

type AreaShort struct {
	Cords       []GeoCords `json:"cords"`
	CreatedDate time.Time  `json:"created_date"`
	Id          int        `json:"id"`
	IsActive    bool       `json:"is_active"`
	Name        string     `json:"name"`
	UpdatedDate time.Time  `json:"updated_date"`
}

type GeoCords struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
