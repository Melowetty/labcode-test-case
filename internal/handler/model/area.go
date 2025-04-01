package model

type CreateAreaRequest struct {
	Cords    []GeoCordsRequest `json:"cords" validate:"required,dive,required"`
	IsActive *bool             `json:"is_active" validate:"required"`
	Name     string            `json:"name" validate:"required"`
}

type UpdateAreaRequest struct {
	Cords    []GeoCordsRequest `json:"cords" validate:"required,dive,required"`
	IsActive *bool             `json:"is_active" validate:"required"`
	Name     string            `json:"name" validate:"required"`
}

type GeoCordsRequest struct {
	Latitude  *float32 `json:"latitude" validate:"required,latitude"`
	Longitude *float32 `json:"longitude" validate:"required,longitude"`
}
