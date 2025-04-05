package service

import (
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/entity"
	"labcode-test-case/internal/utils"
)

func areaDetailedToDetailedDto(area entity.AreaDetailed) dto.AreaDetailed {
	return dto.AreaDetailed{
		Cords:       cordsToDto(area.Cords),
		CreatedDate: area.CreatedDate,
		Id:          area.Id,
		IsActive:    area.IsActive,
		Name:        area.Name,
		UpdatedDate: area.UpdatedDate,
		Cameras:     utils.ConvertArray(area.Cameras, cameraToDto),
	}
}

func areaToDto(area entity.Area) dto.AreaShort {
	return dto.AreaShort{
		Cords:       cordsToDto(area.Cords),
		CreatedDate: area.CreatedDate,
		Id:          area.Id,
		IsActive:    area.IsActive,
		Name:        area.Name,
		UpdatedDate: area.UpdatedDate,
	}
}

func cordsToDto(cords []entity.GeoCords) []dto.GeoCords {
	return utils.ConvertArray(cords, func(value entity.GeoCords) dto.GeoCords {
		return dto.GeoCords{
			Latitude:  value.Latitude,
			Longitude: value.Longitude,
		}
	})
}

func cameraToDto(camera entity.Camera) dto.Camera {
	return dto.Camera{
		Altitude:    camera.Altitude,
		Angle:       camera.Angle,
		AreaId:      camera.AreaId,
		CreatedDate: camera.CreatedDate,
		Id:          camera.Id,
		Ip:          camera.Ip,
		IsActive:    camera.IsActive,
		Latitude:    camera.Latitude,
		Longitude:   camera.Longitude,
		Name:        camera.Name,
		Radius:      camera.Radius,
		SectorAngle: camera.SectorAngle,
		UpdatedDate: camera.UpdatedDate,
	}
}

func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
