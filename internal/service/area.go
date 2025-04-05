package service

import (
	"context"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/entity"
	"labcode-test-case/internal/handler/model"
	"labcode-test-case/internal/utils"
	"time"
)

type AreaStorage interface {
	GetAreas(ctx context.Context) ([]entity.Area, error)
	GetAreaById(ctx context.Context, id int) (entity.AreaDetailed, error)
	SaveArea(ctx context.Context, area entity.Area) (entity.AreaDetailed, error)
	DeleteArea(ctx context.Context, id int) error
}

type AreaService struct {
	areaStorage AreaStorage
}

func NewAreaService(storage AreaStorage) *AreaService {
	service := &AreaService{areaStorage: storage}
	return service
}

func (a *AreaService) CreateArea(ctx context.Context, request model.CreateAreaRequest) (dto.AreaDetailed, error) {
	cords := utils.ConvertArray(request.Cords, func(value model.GeoCordsRequest) entity.GeoCords {
		return entity.GeoCords{
			Latitude:  *value.Latitude,
			Longitude: *value.Longitude,
		}
	})

	area := entity.Area{
		Cords:       cords,
		CreatedDate: time.Time{},
		Id:          0,
		IsActive:    *request.IsActive,
		Name:        request.Name,
		UpdatedDate: time.Time{},
	}
	area.UpdatedDate = time.Now()
	area.CreatedDate = time.Now()

	result, err := a.areaStorage.SaveArea(ctx, area)

	if err != nil {
		return dto.AreaDetailed{}, err
	}

	return areaDetailedToDetailedDto(result), nil
}

func (a *AreaService) UpdateArea(ctx context.Context, areaId int, request model.UpdateAreaRequest) (dto.AreaDetailed, error) {
	cords := utils.ConvertArray(request.Cords, func(value model.GeoCordsRequest) entity.GeoCords {
		return entity.GeoCords{
			Latitude:  *value.Latitude,
			Longitude: *value.Longitude,
		}
	})

	area := entity.Area{
		Cords:       cords,
		CreatedDate: time.Time{},
		Id:          areaId,
		IsActive:    *request.IsActive,
		Name:        request.Name,
		UpdatedDate: time.Time{},
	}
	area.UpdatedDate = time.Now()

	result, err := a.areaStorage.SaveArea(ctx, area)

	if err != nil {
		return dto.AreaDetailed{}, err
	}

	return areaDetailedToDetailedDto(result), nil
}
func (a *AreaService) GetArea(ctx context.Context, areaId int) (dto.AreaDetailed, error) {
	area, err := a.areaStorage.GetAreaById(ctx, areaId)

	if err != nil {
		return dto.AreaDetailed{}, err
	}
	return areaDetailedToDetailedDto(area), nil
}
func (a *AreaService) GetAreas(ctx context.Context) ([]dto.AreaShort, error) {
	areas, err := a.areaStorage.GetAreas(ctx)

	if err != nil {
		return nil, err
	}

	areasDto := utils.ConvertArray(areas, areaToDto)
	return areasDto, nil
}
func (a *AreaService) DeleteArea(ctx context.Context, areaId int) error {
	err := a.areaStorage.DeleteArea(ctx, areaId)
	return err
}
