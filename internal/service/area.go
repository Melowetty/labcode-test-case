package service

import (
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/handler/model"
)

type AreaService struct{}

func (a *AreaService) CreateArea(request model.CreateAreaRequest) (dto.AreaDetailed, error) {
	return dto.AreaDetailed{}, nil
}

func (a *AreaService) UpdateArea(areaId int, request model.UpdateAreaRequest) (dto.AreaDetailed, error) {
	return dto.AreaDetailed{}, nil
}
func (a *AreaService) GetArea(areaId int) (dto.AreaDetailed, error) {
	if areaId == 2 {
		return dto.AreaDetailed{}, dto.AreaNotFoundError.Here().WithUserMessagef("Area by id %d not found", areaId)
	}
	return dto.AreaDetailed{}, nil
}
func (a *AreaService) GetAreas() ([]dto.AreaShort, error) {
	return []dto.AreaShort{}, nil
}
func (a *AreaService) DeleteArea(areaId int) error {
	return nil
}
