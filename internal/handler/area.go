package handler

import (
	"fmt"
	"github.com/ansel1/merry"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/handler/model"
	"net/http"
)

type AreaService interface {
	CreateArea(request model.CreateAreaRequest) (dto.AreaDetailed, error)
	UpdateArea(areaId int, request model.UpdateAreaRequest) (dto.AreaDetailed, error)
	GetArea(areaId int) (dto.AreaDetailed, error)
	GetAreas() ([]dto.AreaShort, error)
	DeleteArea(areaId int) error
}

type AreaHandler struct {
	validate    *validator.Validate
	areaService AreaService
}

const (
	AreaIdValidatorError = "Area id must be integer and greater than 0"
)

func NewAreaHandler(mux *http.ServeMux, validate *validator.Validate, areaService AreaService) *AreaHandler {
	controller := &AreaHandler{validate, areaService}
	controller.initRouter(mux)
	return controller
}

func (a *AreaHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area", a.getAreasInfo)
	mux.HandleFunc("GET /area/{id}", a.getAreaInfo)
	mux.HandleFunc("POST /area", a.createArea)
	mux.HandleFunc("DELETE /area", a.deleteArea)
	mux.HandleFunc("PUT /area/{id}", a.updateArea)
}

// @Summary Update area
// @Tags Зона
// @Param area  body      model.UpdateAreaRequest  true  "Area JSON"
// @Param id  path      int  true  "area id"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area/{id} [put]
func (a *AreaHandler) updateArea(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, AreaIdValidatorError, http.StatusBadRequest)
		return
	}

	var request model.UpdateAreaRequest
	err = parseBody(r, &request)
	if err != nil {
		fmt.Println(err.Error())
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(a.validate, request)
	if validationError != nil {
		fmt.Println(validationError.Error())
		writeValidationErrorsResponse(w, validationError)
		return
	}

	area, err := a.areaService.UpdateArea(areaId, request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	fmt.Fprintf(w, "Area updated + %#v", area)
	w.WriteHeader(http.StatusOK)
}

// @Summary      Delete area by id
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Success      200
// @Router       /area/{id} [delete]
func (a *AreaHandler) deleteArea(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, AreaIdValidatorError, http.StatusBadRequest)
		return
	}

	deleteErr := a.areaService.DeleteArea(areaId)
	if deleteErr != nil {
		processErrorResponse(w, deleteErr)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Create area
// @Tags Зоны
// @Param area  body      model.CreateAreaRequest  true  "Area JSON"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area [post]
func (a *AreaHandler) createArea(w http.ResponseWriter, r *http.Request) {
	var request model.CreateAreaRequest
	err := parseBody(r, &request)
	if err != nil {
		fmt.Println(err.Error())
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(a.validate, request)
	if validationError != nil {
		fmt.Println(validationError.Error())
		writeValidationErrorsResponse(w, validationError)
		return
	}

	area, err := a.areaService.CreateArea(request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	fmt.Fprintf(w, "Area created + %#v", area)
}

// @Summary Get areas
// @Tags Зоны
// @Produce json
// @Success 200 {array} dto.AreaShort
// @Router /area [get]
func (a *AreaHandler) getAreasInfo(w http.ResponseWriter, r *http.Request) {
	areas, err := a.areaService.GetAreas()
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	fmt.Fprintln(w, "Areas get "+string(rune(len(areas))))
}

// @Summary Get area info
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area/{id} [get]
func (a *AreaHandler) getAreaInfo(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(a.validate, r, "id")
	if err != nil {
		writeError(w, AreaIdValidatorError, http.StatusBadRequest)
		return
	}

	area, err := a.areaService.GetArea(areaId)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	fmt.Fprintf(w, "Area get %#v", area)
}

func processErrorResponse(w http.ResponseWriter, err error) {
	if merry.Is(err, dto.BaseError) {
		code := merry.HTTPCode(err)
		userMessage := merry.UserMessage(err)
		writeError(w, userMessage, code)
	} else {
		fmt.Println(err.Error())
		writeInternalServerError(w)
	}
}
