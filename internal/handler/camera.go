package handler

import (
	"fmt"
	"net/http"
)

type CameraHandler struct{}

func NewCameraHandler(mux *http.ServeMux) *CameraHandler {
	controller := &CameraHandler{}
	controller.initRouter(mux)
	return controller
}

func (c *CameraHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area/{id}/camera/{camera_id}", c.getCamera)
	mux.HandleFunc("POST /area/{id}/camera", c.createCamera)
	mux.HandleFunc("DELETE /area/{id}/camera/{camera_id}", c.deleteCamera)
	mux.HandleFunc("PUT /area/{id}/camera/{camera_id}", c.updateCamera)
}

func (c *CameraHandler) updateCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *CameraHandler) deleteCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *CameraHandler) createCamera(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Area created")
}

func (c *CameraHandler) getCamera(w http.ResponseWriter, r *http.Request) {
	areaId := r.PathValue("id")
	cameraId := r.PathValue("camera_id")
	fmt.Fprintln(w, "Camera get by area "+areaId+" and camera "+cameraId)
}
