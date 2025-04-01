package dto

import "time"

type AreaDetailed struct {
	Cords      []GeoCords `json:"cords"`
	CreateDate time.Time  `json:"create_date"`
	Id         int        `json:"id"`
	IsActive   bool       `json:"is_active"`
	Name       string     `json:"name"`
	UpdateDate time.Time  `json:"update_date"`
	Cameras    []Camera   `json:"cameras"`
}

type AreaShort struct {
	Cords      []GeoCords `json:"cords"`
	CreateDate time.Time  `json:"create_date"`
	Id         int        `json:"id"`
	IsActive   bool       `json:"is_active"`
	Name       string     `json:"name"`
	UpdateDate time.Time  `json:"update_date"`
}

type GeoCords struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
