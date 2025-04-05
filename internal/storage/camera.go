package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"labcode-test-case/internal/dto"
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

func (s *CameraStorage) GetCamera(ctx context.Context, cameraId int) (entity.Camera, error) {
	query := `SELECT id, name, altitude, angle, area_id, latitude, longitude, radius, sector_angle, is_active, 
       ip, created_date, updated_date FROM public.camera WHERE id = $1`
	rows, err := s.pool.Query(ctx, query, cameraId)
	if err != nil {
		return entity.Camera{}, fmt.Errorf("failed to get camera: %w", err)
	}
	defer rows.Close()

	camera, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Camera])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Camera{}, dto.CameraNotFoundError.Here().WithUserMessagef("Camera by id %d not found", cameraId)
		}
		return entity.Camera{}, fmt.Errorf("failed collecting rows: %w", err)
	}

	return camera, nil
}

func (s *CameraStorage) SaveCamera(ctx context.Context, camera entity.Camera) (entity.Camera, error) {
	if camera.Id != 0 {
		return s.updateCamera(ctx, camera)
	}

	query := `INSERT INTO public.camera(name, altitude, angle, area_id, latitude, longitude, radius, sector_angle, is_active, ip, created_date, updated_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`

	err := s.pool.QueryRow(ctx, query, camera.Name, camera.Altitude, camera.Angle, camera.AreaId, camera.Latitude, camera.Longitude,
		camera.Radius, camera.SectorAngle, camera.IsActive, camera.Ip, camera.CreatedDate, camera.UpdatedDate).Scan(&camera.Id)

	if err != nil {
		return entity.Camera{}, fmt.Errorf("failed to create camera %w", err)
	}

	return camera, nil
}

func (s *CameraStorage) updateCamera(ctx context.Context, camera entity.Camera) (entity.Camera, error) {
	existsCamera, err := s.GetCamera(ctx, camera.Id)
	if err != nil {
		return entity.Camera{}, err
	}

	query := `UPDATE public.camera SET name=$1, altitude=$2, angle=$3, area_id=$4, latitude=$5, longitude=$6, 
		radius=$7, sector_angle=$8, is_active=$9, ip=$10, updated_date=$11 WHERE id=$12`
	_, err = s.pool.Query(ctx, query, camera.Name, camera.Altitude, camera.Angle, camera.AreaId, camera.Latitude, camera.Longitude, camera.Radius,
		camera.SectorAngle, camera.IsActive, camera.Ip, camera.UpdatedDate, camera.Id)

	if err != nil {
		return entity.Camera{}, fmt.Errorf("failed to update camera %w", err)
	}

	camera.CreatedDate = existsCamera.CreatedDate

	return camera, nil
}

func (s *CameraStorage) DeleteCamera(ctx context.Context, cameraId int) error {
	query := "DELETE FROM camera WHERE id = $1"
	res, err := s.pool.Exec(ctx, query, cameraId)

	if err != nil {
		return fmt.Errorf("failed deleting camera: %w", err)
	}

	if res.RowsAffected() == 0 {
		return dto.CameraNotFoundError.Here().WithUserMessagef("Camera by id %d not found", cameraId)
	}

	return nil
}

func (s *CameraStorage) GetCameras(ctx context.Context, areaId int) ([]entity.Camera, error) {
	query := `SELECT id, name, altitude, angle, area_id, latitude, longitude, radius, sector_angle, is_active, 
       ip, created_date, updated_date FROM public.camera WHERE area_id = $1`
	rows, err := s.pool.Query(ctx, query, areaId)
	if err != nil {
		return nil, fmt.Errorf("failed to get cameras: %w", err)
	}
	defer rows.Close()

	cameras, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Camera])
	if err != nil {
		return nil, fmt.Errorf("failed collecting rows: %w", err)
	}

	return cameras, nil
}
