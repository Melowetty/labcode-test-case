package model

type CreateCameraRequest struct {
	Altitude    *float32 `json:"altitude" validate:"required"`
	Angle       *float32 `json:"angle" validate:"required,gte=0"`
	Ip          string   `json:"ip" validate:"required,ip"`
	IsActive    *bool    `json:"is_active" validate:"required"`
	Latitude    *float32 `json:"latitude" validate:"required,latitude"`
	Longitude   *float32 `json:"longitude" validate:"required,longitude"`
	Name        string   `json:"name" validate:"required"`
	Radius      *float32 `json:"radius" validate:"required,gt=0"`
	SectorAngle *float32 `json:"sector_angle" validate:"required,gte=0"`
}

type UpdateCameraRequest struct {
	Altitude    *float32 `json:"altitude" validate:"required"`
	Angle       *float32 `json:"angle" validate:"required,gte=0"`
	Ip          string   `json:"ip" validate:"required,ip"`
	IsActive    *bool    `json:"is_active" validate:"required"`
	Latitude    *float32 `json:"latitude" validate:"required,latitude"`
	Longitude   *float32 `json:"longitude" validate:"required,longitude"`
	Name        string   `json:"name" validate:"required"`
	Radius      *float32 `json:"radius" validate:"required,gt=0"`
	SectorAngle *float32 `json:"sector_angle" validate:"required,gte=0"`
}
