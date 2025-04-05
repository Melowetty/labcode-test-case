package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"labcode-test-case/internal/entity"
)

type CameraStorage struct {
	pool *pgxpool.Pool
}

func NewCameraStorage(pool *pgxpool.Pool) *CameraStorage {
	return &CameraStorage{
		pool: pool,
	}
}

func (s *CameraStorage) GetCameras(ctx context.Context, areaId int) ([]entity.Camera, error) {
	return []entity.Camera{}, nil
}
