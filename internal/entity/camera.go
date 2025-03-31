package entity

type Camera struct {
	Altitude    int    `json:"altitude"`
	Angle       int    `json:"angle"`
	AreaId      int    `json:"area_id"`
	CreateBy    int    `json:"create_by"`
	CreateDate  string `json:"create_date"`
	Id          int    `json:"id"`
	Ip          string `json:"ip"`
	IsActive    bool   `json:"is_active"`
	Latitude    int    `json:"latitude"`
	Longitude   int    `json:"longitude"`
	Name        string `json:"name"`
	Radius      int    `json:"radius"`
	SectorAngle int    `json:"sector_angle"`
	UpdateBy    int    `json:"update_by"`
	UpdateDate  string `json:"update_date"`
}
