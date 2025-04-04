package service

import (
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/entity"
	"labcode-test-case/internal/handler/model"
	"labcode-test-case/internal/utils"
	"time"
)

type AreaStorage interface {
	GetAreas() ([]entity.Area, error)
	GetAreaById(id int) (entity.AreaDetailed, error)
	SaveArea(area entity.Area) (entity.AreaDetailed, error)
	DeleteArea(id int) error
}

type AreaService struct {
	areaStorage AreaStorage
}

func NewAreaService(storage AreaStorage) *AreaService {
	service := &AreaService{areaStorage: storage}
	return service
}

func (a *AreaService) CreateArea(request model.CreateAreaRequest) (dto.AreaDetailed, error) {
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

	result, err := a.areaStorage.SaveArea(area)

	return areaDetailedToDetailedDto(result), err
}

func (a *AreaService) UpdateArea(areaId int, request model.UpdateAreaRequest) (dto.AreaDetailed, error) {
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

	result, err := a.areaStorage.SaveArea(area)

	return areaDetailedToDetailedDto(result), err
}
func (a *AreaService) GetArea(areaId int) (dto.AreaDetailed, error) {
	area, err := a.areaStorage.GetAreaById(areaId)
	return areaDetailedToDetailedDto(area), err
}
func (a *AreaService) GetAreas() ([]dto.AreaShort, error) {
	areas, err := a.areaStorage.GetAreas()

	areasDto := utils.ConvertArray(areas, areaToDto)
	return areasDto, err
}
func (a *AreaService) DeleteArea(areaId int) error {
	err := a.areaStorage.DeleteArea(areaId)
	return err
}
