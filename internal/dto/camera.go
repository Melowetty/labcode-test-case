package dto

import "time"

type Camera struct {
	Altitude    float32   `json:"altitude"`
	Angle       float32   `json:"angle"`
	AreaId      int       `json:"area_id"`
	CreatedDate time.Time `json:"created_date"`
	Id          int       `json:"id"`
	Ip          string    `json:"ip"`
	IsActive    bool      `json:"is_active"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	Name        string    `json:"name"`
	Radius      float32   `json:"radius"`
	SectorAngle float32   `json:"sector_angle"`
	UpdatedDate time.Time `json:"updated_date"`
}
