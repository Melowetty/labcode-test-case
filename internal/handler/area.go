package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/handler/model"
	"net/http"
)

type AreaServiceInterface interface {
	CreateArea(ctx context.Context, request model.CreateAreaRequest) (dto.AreaDetailed, error)
	UpdateArea(ctx context.Context, areaId int, request model.UpdateAreaRequest) (dto.AreaDetailed, error)
	GetArea(ctx context.Context, areaId int) (dto.AreaDetailed, error)
	GetAreas(ctx context.Context) ([]dto.AreaShort, error)
	DeleteArea(ctx context.Context, areaId int) error
}

type AreaHandler struct {
	validate    *validator.Validate
	areaService AreaServiceInterface
}

const (
	areaIdValidatorError = "area id must be integer and greater than 0"
)

func NewAreaHandler(mux *http.ServeMux, validate *validator.Validate, areaService AreaServiceInterface) *AreaHandler {
	controller := &AreaHandler{validate, areaService}
	controller.initRouter(mux)
	return controller
}

func (a *AreaHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area", a.GetAreasInfo)
	mux.HandleFunc("GET /area/{id}", a.GetAreaInfo)
	mux.HandleFunc("POST /area", a.CreateArea)
	mux.HandleFunc("DELETE /area/{id}", a.DeleteArea)
	mux.HandleFunc("PUT /area/{id}", a.UpdateArea)
}

// UpdateArea @Summary Update area
// @Tags Зона
// @Param area  body      model.UpdateAreaRequest  true  "Area JSON"
// @Param id  path      int  true  "area id"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area/{id} [put]
func (a *AreaHandler) UpdateArea(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, areaIdValidatorError, http.StatusBadRequest)
		return
	}

	var request model.UpdateAreaRequest
	err = parseBody(r, &request)
	if err != nil {
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(a.validate, request)
	if validationError != nil {
		writeValidationErrorsResponse(w, validationError)
		return
	}

	ctx := r.Context()
	area, err := a.areaService.UpdateArea(ctx, areaId, request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, area)
}

// DeleteArea @Summary      Delete area by id
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Success      200
// @Router       /area/{id} [delete]
func (a *AreaHandler) DeleteArea(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, areaIdValidatorError, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	deleteErr := a.areaService.DeleteArea(ctx, areaId)
	if deleteErr != nil {
		processErrorResponse(w, deleteErr)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CreateArea @Summary Create area
// @Tags Зоны
// @Param area  body      model.CreateAreaRequest  true  "Area JSON"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area [post]
func (a *AreaHandler) CreateArea(w http.ResponseWriter, r *http.Request) {
	var request model.CreateAreaRequest
	err := parseBody(r, &request)
	if err != nil {
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(a.validate, request)
	if validationError != nil {
		writeValidationErrorsResponse(w, validationError)
		return
	}

	ctx := r.Context()
	area, err := a.areaService.CreateArea(ctx, request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, area)
}

// GetAreasInfo @Summary Get areas
// @Tags Зоны
// @Produce json
// @Success 200 {array} dto.AreaShort
// @Router /area [get]
func (a *AreaHandler) GetAreasInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	areas, err := a.areaService.GetAreas(ctx)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, areas)
}

// GetAreaInfo @Summary Get area info
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area/{id} [get]
func (a *AreaHandler) GetAreaInfo(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, areaIdValidatorError, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	area, err := a.areaService.GetArea(ctx, areaId)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, area)
}
