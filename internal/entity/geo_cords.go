package entity

type GeoCords struct {
	AreaId    int     `db:"area_id"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`
}
