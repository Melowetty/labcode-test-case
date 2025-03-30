package controller

import (
	"fmt"
	"net/http"
)

type CameraController struct{}

func NewCameraController() *CameraController {
	controller := &CameraController{}
	controller.initRouter()
	return controller
}

func (c *CameraController) initRouter() {
	http.HandleFunc("/camera", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			c.getCameraInfo(w, r)
		case "POST":
			c.createCamera(w, r)
		case "DELETE":
			c.deleteCamera(w, r)
		case "PUT":
			c.updateCamera(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func (c *CameraController) updateCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *CameraController) deleteCamera(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c *CameraController) createCamera(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func (c *CameraController) getCameraInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
