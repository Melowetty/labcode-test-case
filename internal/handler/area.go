package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/handler/model"
	"net/http"
)

type AreaHandler struct {
	validate *validator.Validate
}

func NewAreaHandler(mux *http.ServeMux, validate *validator.Validate) *AreaHandler {
	controller := &AreaHandler{validate}
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
// @Router /area [put]
func (a *AreaHandler) updateArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary      Delete area by id
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Success      200
// @Router       /area/{id} [delete]
func (a *AreaHandler) deleteArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary Create area
// @Tags Зоны
// @Param area  body      model.CreateAreaRequest  true  "Area JSON"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area [post]
func (a *AreaHandler) createArea(w http.ResponseWriter, r *http.Request) {
	var area model.CreateAreaRequest
	err := parseBody(r, &area)
	if err != nil {
		fmt.Println(err.Error())
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(a.validate, area)
	if validationError != nil {
		fmt.Println(validationError.Error())
		writeValidationErrorsResponse(w, validationError)
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
	fmt.Fprintln(w, "Areas get")
}

// @Summary Get area info
// @Tags Зоны
// @Param        id  path      int  true  "area id"
// @Produce json
// @Success 200 {object} dto.AreaDetailed
// @Router /area/{id} [get]
func (a *AreaHandler) getAreaInfo(w http.ResponseWriter, r *http.Request) {
	areaId := r.PathValue("id")
	fmt.Fprintln(w, "Area get "+areaId)
}
