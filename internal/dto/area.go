package dto

import "time"

type AreaDetailed struct {
	Cords       []GeoCords `json:"cords"`
	CreatedDate time.Time  `json:"created_date" example:"2025-04-06T03:42:06.553269Z"`
	Id          int        `json:"id" example:"123"`
	IsActive    bool       `json:"is_active" example:"true"`
	Name        string     `json:"name" example:"Perm"`
	UpdatedDate time.Time  `json:"updated_date" example:"2025-04-06T03:42:06.553269Z"`
	Cameras     []Camera   `json:"cameras"`
}

type AreaShort struct {
	Cords       []GeoCords `json:"cords"`
	CreatedDate time.Time  `json:"created_date" example:"2025-04-06T03:42:06.553269Z"`
	Id          int        `json:"id" example:"123"`
	IsActive    bool       `json:"is_active" example:"true"`
	Name        string     `json:"name" example:"Perm"`
	UpdatedDate time.Time  `json:"updated_date" example:"2025-04-06T03:42:06.553269Z"`
}

type GeoCords struct {
	Latitude  float32 `json:"latitude" example=:"-10.31342"`
	Longitude float32 `json:"longitude" example=:"50.12354"`
}
