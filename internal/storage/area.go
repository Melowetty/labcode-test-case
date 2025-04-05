package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/entity"
	"labcode-test-case/internal/utils"
	"time"
)

type CameraStorageInterface interface {
	GetCameras(ctx context.Context, areaId int) ([]entity.Camera, error)
}

type AreaStorage struct {
	pool          *pgxpool.Pool
	cameraStorage CameraStorageInterface
}

type clearArea struct {
	CreatedDate time.Time `db:"created_date"`
	Id          int       `db:"id"`
	IsActive    bool      `db:"is_active"`
	Name        string    `db:"name"`
	UpdatedDate time.Time `db:"updated_date"`
}

func NewAreaStorage(pool *pgxpool.Pool, cameraStorage CameraStorageInterface) *AreaStorage {
	return &AreaStorage{
		pool:          pool,
		cameraStorage: cameraStorage,
	}
}

func (s *AreaStorage) GetAreas(ctx context.Context) ([]entity.Area, error) {
	areas, err := s.getAreasEntityWithoutCoords(ctx)
	if err != nil {
		return nil, err
	}
	areasId := utils.ConvertArray(areas, func(value clearArea) int {
		return value.Id
	})

	cords, err := s.getCordsByAreaId(ctx, areasId...)
	if err != nil {
		return nil, err
	}

	cordsByArea := make(map[int][]entity.GeoCords)
	for _, c := range cords {
		cordsByArea[c.AreaId] = append(cordsByArea[c.AreaId], c)
	}

	result := make([]entity.Area, 0, len(areas))
	for _, area := range areas {
		newArea := clearAreaToEntity(area, cordsByArea[area.Id])
		result = append(result, newArea)
	}

	return result, nil
}

func clearAreaToEntity(area clearArea, cords []entity.GeoCords) entity.Area {
	return entity.Area{
		Id:          area.Id,
		Name:        area.Name,
		IsActive:    area.IsActive,
		CreatedDate: area.CreatedDate,
		UpdatedDate: area.UpdatedDate,
		Cords:       cords,
	}
}

func (s *AreaStorage) getAreasEntityWithoutCoords(ctx context.Context) ([]clearArea, error) {
	query := "SELECT id, name, is_active, created_date, updated_date FROM area"
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get areas: %w", err)
	}
	defer rows.Close()

	areas, err := pgx.CollectRows(rows, pgx.RowToStructByName[clearArea])

	if err != nil {
		return nil, fmt.Errorf("failed collecting rows: %w", err)
	}

	return areas, nil
}

func (s *AreaStorage) getCordsByAreaId(ctx context.Context, areaIds ...int) ([]entity.GeoCords, error) {
	query := "SELECT area_id, latitude, longitude  FROM cord WHERE area_id = any($1)"
	rows, err := s.pool.Query(ctx, query, areaIds)
	if err != nil {
		return nil, fmt.Errorf("failed to get cords: %w", err)
	}
	defer rows.Close()

	cords, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.GeoCords])

	if err != nil {
		return nil, fmt.Errorf("failed collecting rows: %w", err)
	}

	return cords, nil
}

func (s *AreaStorage) GetAreaById(ctx context.Context, id int) (entity.AreaDetailed, error) {
	query := "SELECT id, name, is_active, created_date, updated_date FROM area WHERE id = $1"
	rows, err := s.pool.Query(ctx, query, id)
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to get area: %w", err)
	}
	defer rows.Close()

	area, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[clearArea])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.AreaDetailed{}, dto.AreaNotFoundError.Here().WithUserMessagef("Area by id %d not found", id)
		}
		return entity.AreaDetailed{}, fmt.Errorf("failed collecting rows: %w", err)
	}

	cords, err := s.getCordsByAreaId(ctx, id)
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to get cords: %w", err)
	}

	cameras, err := s.cameraStorage.GetCameras(ctx, id)
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to get cameras: %w", err)
	}

	result := entity.AreaDetailed{
		Area:    clearAreaToEntity(area, cords),
		Cameras: cameras,
	}
	return result, nil
}

func (s *AreaStorage) SaveArea(ctx context.Context, area entity.Area) (entity.AreaDetailed, error) {
	if area.Id != 0 {
		return s.updateArea(ctx, area)
	}
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to begin tx: %w", err)
	}

	query := "INSERT INTO area (name, is_active, created_date, updated_date) VALUES ($1, $2, $3, $4) RETURNING id"
	err = tx.QueryRow(ctx, query, area.Name, area.IsActive, area.CreatedDate, area.UpdatedDate).Scan(&area.Id)
	if err != nil {
		tx.Rollback(ctx)
		return entity.AreaDetailed{}, fmt.Errorf("failed to insert area: %w", err)
	}

	batch := pgx.Batch{}
	query = "INSERT INTO cord (area_id, longitude, latitude) VALUES ($1, $2, $3)"
	for _, cord := range area.Cords {
		batch.Queue(query, area.Id, cord.Longitude, cord.Latitude)
	}
	res := tx.SendBatch(ctx, &batch)
	if err := res.Close(); err != nil {
		tx.Rollback(ctx)
		return entity.AreaDetailed{}, fmt.Errorf("failed to save cords: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to commit tx: %w", err)
	}

	detailedEntity := entity.AreaDetailed{Area: area, Cameras: []entity.Camera{}}
	return detailedEntity, nil
}

func (s *AreaStorage) updateArea(ctx context.Context, area entity.Area) (entity.AreaDetailed, error) {
	existsArea, err := s.GetAreaById(ctx, area.Id)
	if err != nil {
		return entity.AreaDetailed{}, err
	}

	area.CreatedDate = existsArea.CreatedDate

	batch := pgx.Batch{}
	query := "UPDATE area SET name = $1, is_active = $2, updated_date = $3 WHERE id = $4"
	batch.Queue(query, area.Name, area.IsActive, area.UpdatedDate, area.Id)

	cords, err := s.getCordsByAreaId(ctx, area.Id)
	if err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to get cords: %w", err)
	}

	existsCordsFromDb := map[string]entity.GeoCords{}
	for _, cord := range cords {
		existsCordsFromDb[fmt.Sprintf("%.6f %.6f", cord.Latitude, cord.Longitude)] = cord
	}
	existsCordsFromRequest := map[string]entity.GeoCords{}
	for _, cord := range area.Cords {
		existsCordsFromRequest[fmt.Sprintf("%.6f %.6f", cord.Latitude, cord.Longitude)] = cord
	}

	deleteQuery := "DELETE FROM cord WHERE area_id = $1 AND latitude = $2 AND longitude = $3"
	for key, cord := range existsCordsFromDb {
		if _, ok := existsCordsFromRequest[key]; !ok {
			batch.Queue(deleteQuery, area.Id, cord.Latitude, cord.Longitude)
		}
	}

	createQuery := "INSERT INTO cord (area_id, longitude, latitude) VALUES ($1, $2, $3)"
	for key, cord := range existsCordsFromRequest {
		if _, ok := existsCordsFromDb[key]; !ok {
			batch.Queue(createQuery, area.Id, cord.Longitude, cord.Latitude)
		}
	}

	res := s.pool.SendBatch(ctx, &batch)
	if err := res.Close(); err != nil {
		return entity.AreaDetailed{}, fmt.Errorf("failed to batch query: %w", err)
	}

	return entity.AreaDetailed{
		Area:    area,
		Cameras: existsArea.Cameras,
	}, nil
}

func (s *AreaStorage) DeleteArea(ctx context.Context, id int) error {
	query := "DELETE FROM area WHERE id = $1"
	res, err := s.pool.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("failed deleting area: %w", err)
	}

	if res.RowsAffected() == 0 {
		return dto.AreaNotFoundError.Here().WithUserMessagef("Area by id %d not found", id)
	}

	return nil
}
