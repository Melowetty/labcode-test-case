package entity

import "time"

type Camera struct {
	Altitude    float32   `db:"altitude"`
	Angle       float32   `db:"angle"`
	AreaId      int       `db:"area_id"`
	CreatedDate time.Time `db:"created_date"`
	Id          int       `db:"id"`
	Ip          string    `db:"ip"`
	IsActive    bool      `db:"is_active"`
	Latitude    float32   `db:"latitude"`
	Longitude   float32   `db:"longitude"`
	Name        string    `db:"name"`
	Radius      float32   `db:"radius"`
	SectorAngle float32   `db:"sector_angle"`
	UpdatedDate time.Time `db:"updated_date"`
}
