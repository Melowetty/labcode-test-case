package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/handler/model"
	"net/http"
)

type CameraHandler struct {
	validate *validator.Validate
}

func NewCameraHandler(mux *http.ServeMux, validate *validator.Validate) *CameraHandler {
	controller := &CameraHandler{validate}
	controller.initRouter(mux)
	return controller
}

func (c *CameraHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area/{id}/camera/{camera_id}", c.getCamera)
	mux.HandleFunc("POST /area/{id}/camera", c.createCamera)
	mux.HandleFunc("DELETE /area/{id}/camera/{camera_id}", c.deleteCamera)
	mux.HandleFunc("PUT /area/{id}/camera/{camera_id}", c.updateCamera)
}

// @Summary Update camera
// @Tags Камеры
// @Param area  body model.UpdateCameraRequest  true  "Camera JSON"
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Produce json
// @Success 200 {object} dto.Camera
// @Router /area/{area_id}/camera/{camera_id} [put]
func (c *CameraHandler) updateCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete camera
// @Tags Камеры
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Success 200
// @Router /area/{area_id}/{camera_id} [delete]
func (c *CameraHandler) deleteCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary Create camera
// @Tags Камеры
// @Param area  body model.CreateCameraRequest  true  "Camera JSON"
// @Param area_id  path int  true  "area id"
// @Produce json
// @Success 200 {object} dto.Camera
// @Router /area/{area_id}/camera [post]
func (c *CameraHandler) createCamera(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(c.validate, r, "id")
	if err != nil {
		writeError(w, "Area id must be integer and greater than 0", http.StatusBadRequest)
		return
	}

	var camera model.CreateCameraRequest
	err = parseBody(r, &camera)
	if err != nil {
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(c.validate, camera)
	if validationError != nil {
		writeValidationErrorsResponse(w, validationError)
		return
	}

	fmt.Fprintf(w, "Camera for area %d created %#v", areaId, camera)
}

// @Summary Get camera
// @Tags Камеры
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Produce json
// @Success 200 {object} dto.Camera
// @Router /area/{area_id}/camera/{camera_id} [get]
func (c *CameraHandler) getCamera(w http.ResponseWriter, r *http.Request) {
	areaId := r.PathValue("id")
	cameraId := r.PathValue("camera_id")
	fmt.Fprintln(w, "Camera get by area "+areaId+" and camera "+cameraId)
}
