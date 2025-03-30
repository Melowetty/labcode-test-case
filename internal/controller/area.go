package controller

import (
	"fmt"
	"net/http"
)

type AreaController struct{}

func NewAreaController() *AreaController {
	controller := &AreaController{}
	controller.initRouter()
	return controller
}

func (a *AreaController) initRouter() {
	http.HandleFunc("/area", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			a.getAreaInfo(w, r)
		case "POST":
			a.createArea(w, r)
		case "DELETE":
			a.deleteArea(w, r)
		case "PUT":
			a.updateArea(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func (a *AreaController) updateArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *AreaController) deleteArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *AreaController) createArea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func (a *AreaController) getAreaInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
