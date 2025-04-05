package dto

import "time"

type Camera struct {
	Altitude    float32   `json:"altitude" example:"27.8"`
	Angle       float32   `json:"angle" example:"30.5"`
	AreaId      int       `json:"area_id" example:"1"`
	CreatedDate time.Time `json:"created_date" example:"2025-04-06T03:42:06.553269Z"`
	Id          int       `json:"id" example:"1"`
	Ip          string    `json:"ip" example:"127.0.0.1"`
	IsActive    bool      `json:"is_active" example:"true"`
	Latitude    float32   `json:"latitude" example:"27.123454"`
	Longitude   float32   `json:"longitude" example:"30.123456"`
	Name        string    `json:"name" example:"Enter"`
	Radius      float32   `json:"radius" example:"10.5"`
	SectorAngle float32   `json:"sector_angle" example:"15.25"`
	UpdatedDate time.Time `json:"updated_date" example:"2025-04-06T03:42:06.553269Z"`
}
