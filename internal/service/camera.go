package service

import (
	"context"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/entity"
	"labcode-test-case/internal/handler/model"
	"time"
)

type CameraStorage interface {
	GetCamera(ctx context.Context, cameraId int) (entity.Camera, error)
	SaveCamera(ctx context.Context, camera entity.Camera) (entity.Camera, error)
	DeleteCamera(ctx context.Context, cameraId int) error
}

type AreaStorageInterface interface {
	GetAreaCords(ctx context.Context, areaId int) ([]entity.GeoCords, error)
}

type CameraService struct {
	cameraStorage CameraStorage
	areaStorage   AreaStorageInterface
}

func NewCameraService(storage CameraStorage, areaStorage AreaStorageInterface) *CameraService {
	service := &CameraService{cameraStorage: storage, areaStorage: areaStorage}
	return service
}

func (c *CameraService) GetCamera(ctx context.Context, areaId int, cameraId int) (dto.Camera, error) {
	camera, err := c.cameraStorage.GetCamera(ctx, cameraId)
	if err != nil {
		return dto.Camera{}, err
	}
	return cameraToDto(camera), nil
}

func (c *CameraService) CreateCamera(ctx context.Context, areaId int, camera model.CreateCameraRequest) (dto.Camera, error) {
	areaCords, err := c.areaStorage.GetAreaCords(ctx, areaId)

	if err != nil {
		return dto.Camera{}, err
	}

	latitude := camera.Latitude
	longitude := camera.Longitude
	if !checkIntersection(*latitude, *longitude, areaCords) {
		return dto.Camera{}, dto.CameraNotInAreaError
	}

	entity := entity.Camera{
		Altitude:    *camera.Altitude,
		Angle:       *camera.Angle,
		AreaId:      areaId,
		CreatedDate: time.Now(),
		Id:          0,
		Ip:          camera.Ip,
		IsActive:    *camera.IsActive,
		Latitude:    *camera.Latitude,
		Longitude:   *camera.Longitude,
		Name:        camera.Name,
		Radius:      *camera.Radius,
		SectorAngle: *camera.SectorAngle,
		UpdatedDate: time.Now(),
	}

	res, err := c.cameraStorage.SaveCamera(ctx, entity)
	if err != nil {
		return dto.Camera{}, err
	}
	return cameraToDto(res), nil
}

func (c *CameraService) UpdateCamera(ctx context.Context, areaId int, cameraId int, camera model.UpdateCameraRequest) (dto.Camera, error) {
	entity := entity.Camera{
		Altitude:    *camera.Altitude,
		Angle:       *camera.Angle,
		AreaId:      areaId,
		CreatedDate: time.Now(),
		Id:          cameraId,
		Ip:          camera.Ip,
		IsActive:    *camera.IsActive,
		Latitude:    *camera.Latitude,
		Longitude:   *camera.Longitude,
		Name:        camera.Name,
		Radius:      *camera.Radius,
		SectorAngle: *camera.SectorAngle,
		UpdatedDate: time.Now(),
	}

	res, err := c.cameraStorage.SaveCamera(ctx, entity)
	if err != nil {
		return dto.Camera{}, err
	}
	return cameraToDto(res), nil
}

func (c *CameraService) DeleteCamera(ctx context.Context, areaId int, cameraId int) error {
	return c.cameraStorage.DeleteCamera(ctx, cameraId)
}

func checkIntersection(latitude, longitude float32, polygon []entity.GeoCords) bool {
	numVertices := len(polygon)
	inside := false

	p1 := polygon[0]
	for i := 1; i <= numVertices; i++ {
		p2 := polygon[i%numVertices]

		if isPointOnLineSegment(latitude, longitude, p1, p2) {
			return true
		}

		if longitude >= min(p1.Longitude, p2.Longitude) && longitude <= max(p1.Longitude, p2.Longitude) {
			if latitude <= max(p1.Latitude, p2.Latitude) {
				intersection := (longitude-p1.Longitude)*(p2.Latitude-p1.Latitude)/(p2.Longitude-p1.Longitude) + p1.Latitude
				if p1.Latitude == p2.Latitude || latitude < intersection {
					inside = !inside
				}
			}
		}
		p1 = p2
	}
	return inside
}

func isPointOnLineSegment(latitude, longitude float32, p1, p2 entity.GeoCords) bool {
	if longitude <= max(p1.Longitude, p2.Longitude) && longitude >= min(p1.Longitude, p2.Longitude) &&
		latitude <= max(p1.Latitude, p2.Latitude) && latitude >= min(p1.Latitude, p2.Latitude) {

		slope1 := (p2.Latitude - p1.Latitude) * (longitude - p1.Longitude)
		slope2 := (latitude - p1.Latitude) * (p2.Longitude - p1.Longitude)

		return slope1 == slope2
	}
	return false
}
