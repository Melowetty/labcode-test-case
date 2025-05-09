package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pion/webrtc/v3"
	"labcode-test-case/internal/dto"
	"labcode-test-case/internal/handler/model"
	"net/http"
)

type CameraServiceInterface interface {
	GetCamera(ctx context.Context, areaId int, cameraId int) (dto.Camera, error)
	CreateCamera(ctx context.Context, areaId int, camera model.CreateCameraRequest) (dto.Camera, error)
	UpdateCamera(ctx context.Context, areaId int, cameraId int, camera model.UpdateCameraRequest) (dto.Camera, error)
	DeleteCamera(ctx context.Context, areaId int, cameraId int) error
}

type CameraStreamServiceInterface interface {
	GetCameraStream(areaId int, cameraId int, offer webrtc.SessionDescription) (webrtc.SessionDescription, error)
}

type CameraHandler struct {
	validate            *validator.Validate
	cameraService       CameraServiceInterface
	cameraStreamService CameraStreamServiceInterface
}

const (
	cameraIdValidatorError = "camera id must be integer and greater than 0"
)

func NewCameraHandler(mux *http.ServeMux, validate *validator.Validate, cameraService CameraServiceInterface, cameraStreamService CameraStreamServiceInterface) *CameraHandler {
	controller := &CameraHandler{validate: validate, cameraService: cameraService, cameraStreamService: cameraStreamService}
	controller.initRouter(mux)
	return controller
}

func (c *CameraHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area/{id}/camera/{camera_id}", c.GetCamera)
	mux.HandleFunc("POST /area/{id}/camera", c.CreateCamera)
	mux.HandleFunc("DELETE /area/{id}/camera/{camera_id}", c.DeleteCamera)
	mux.HandleFunc("PUT /area/{id}/camera/{camera_id}", c.UpdateCamera)
	mux.HandleFunc("POST /area/{id}/camera/{camera_id}/stream", c.GetCameraStream)
}

// UpdateCamera @Summary Update camera
// @Tags Камеры
// @Param area  body model.UpdateCameraRequest  true  "Camera JSON"
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Produce json
// @Success 200 {object} dto.Camera
// @failure 400 {object} model.CustomError "Wrong query params"
// @failure 400 {object} model.CustomError "Bad body scheme"
// @failure 400 {object} model.CustomError "Validation body error"
// @failure 404 {object} model.CustomError "Entity not found"
// @Router /area/{area_id}/camera/{camera_id} [put]
func (c *CameraHandler) UpdateCamera(w http.ResponseWriter, r *http.Request) {
	areaId, cameraId, err := parseAreaIdAndCameraId(c.validate, w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request model.UpdateCameraRequest
	err = parseBody(r, &request)
	if err != nil {
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(c.validate, request)
	if validationError != nil {
		writeValidationErrorsResponse(w, validationError)
		return
	}

	ctx := r.Context()
	camera, err := c.cameraService.UpdateCamera(ctx, areaId, cameraId, request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, camera)
}

// DeleteCamera @Summary Delete camera
// @Tags Камеры
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Success 200
// @failure 400 {object} model.CustomError "Wrong query params"
// @failure 404 {object} model.CustomError "Entity not found"
// @Router /area/{area_id}/{camera_id} [delete]
func (c *CameraHandler) DeleteCamera(w http.ResponseWriter, r *http.Request) {
	areaId, cameraId, err := parseAreaIdAndCameraId(c.validate, w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	deleteErr := c.cameraService.DeleteCamera(ctx, areaId, cameraId)
	if deleteErr != nil {
		processErrorResponse(w, deleteErr)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CreateCamera @Summary Create camera
// @Tags Камеры
// @Param area  body model.CreateCameraRequest  true  "Camera JSON"
// @Param area_id  path int  true  "area id"
// @Produce json
// @Success 200 {object} dto.Camera
// @failure 400 {object} model.CustomError "Bad body scheme"
// @failure 400 {object} model.CustomError "Validation body error"
// @Router /area/{area_id}/camera [post]
func (c *CameraHandler) CreateCamera(w http.ResponseWriter, r *http.Request) {
	areaId, err := parsePathValueAsInt(c.validate, r, "id")
	if err != nil {
		writeError(w, areaIdValidatorError, http.StatusBadRequest)
		return
	}

	var request model.CreateCameraRequest
	err = parseBody(r, &request)
	if err != nil {
		writeError(w, "Bad body scheme", http.StatusBadRequest)
		return
	}

	validationError := validateBody(c.validate, request)
	if validationError != nil {
		writeValidationErrorsResponse(w, validationError)
		return
	}

	ctx := r.Context()
	camera, err := c.cameraService.CreateCamera(ctx, areaId, request)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, camera)
}

// GetCamera @Summary Get camera
// @Tags Камеры
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Produce json
// @Success 200 {object} dto.Camera
// @failure 400 {object} model.CustomError "Wrong query params"
// @failure 404 {object} model.CustomError "Entity not found"
// @Router /area/{area_id}/camera/{camera_id} [get]
func (c *CameraHandler) GetCamera(w http.ResponseWriter, r *http.Request) {
	areaId, cameraId, err := parseAreaIdAndCameraId(c.validate, w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	camera, err := c.cameraService.GetCamera(ctx, areaId, cameraId)
	if err != nil {
		processErrorResponse(w, err)
		return
	}
	writeOkJsonResponse(w, camera)
}

// GetCameraStream @Summary Get camera stream
// @Tags Камеры
// @Param area_id  path int  true  "area id"
// @Param camera_id  path int  true  "camera id"
// @Router /area/{area_id}/camera/{camera_id}/stream [post]
func (c *CameraHandler) GetCameraStream(w http.ResponseWriter, r *http.Request) {
	areaId, cameraId, err := parseAreaIdAndCameraId(c.validate, w, r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	offer := webrtc.SessionDescription{}
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		processErrorResponse(w, err)
		return
	}

	description, err := c.cameraStreamService.GetCameraStream(areaId, cameraId, offer)
	if err != nil {
		processErrorResponse(w, err)
		return
	}

	writeOkJsonResponse(w, description)
}

func parseAreaIdAndCameraId(validate *validator.Validate, w http.ResponseWriter, r *http.Request) (int, int, error) {
	areaId, err := parsePathValueAsInt(validate, r, "id")
	if err != nil {
		return 0, 0, fmt.Errorf(areaIdValidatorError)
	}
	cameraId, err := parsePathValueAsInt(validate, r, "camera_id")
	if err != nil {
		writeError(w, cameraIdValidatorError, http.StatusBadRequest)
		return 0, 0, fmt.Errorf(cameraIdValidatorError)
	}
	return areaId, cameraId, nil
}
