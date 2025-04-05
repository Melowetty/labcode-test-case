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

type CameraService struct {
	cameraStorage CameraStorage
}

func NewCameraService(storage CameraStorage) *CameraService {
	service := &CameraService{cameraStorage: storage}
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
