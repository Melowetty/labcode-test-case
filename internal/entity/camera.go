package entity

import "time"

type Camera struct {
	Altitude    float32
	Angle       float32
	AreaId      int
	CreatedDate time.Time
	Id          int
	Ip          string
	IsActive    bool
	Latitude    float32
	Longitude   float32
	Name        string
	Radius      float32
	SectorAngle float32
	UpdatedDate time.Time
}
