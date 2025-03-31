package handler

import (
	"fmt"
	"net/http"
)

type AreaHandler struct{}

func NewAreaHandler(mux *http.ServeMux) *AreaHandler {
	controller := &AreaHandler{}
	controller.initRouter(mux)
	return controller
}

func (a *AreaHandler) initRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /area", a.getAreasInfo)
	mux.HandleFunc("GET /area/{id}", a.getAreaInfo)
	mux.HandleFunc("POST /area", a.createArea)
	mux.HandleFunc("DELETE /area", a.deleteArea)
	mux.HandleFunc("PUT /area", a.updateArea)
}

func (a *AreaHandler) updateArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *AreaHandler) deleteArea(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *AreaHandler) createArea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Area created")
}

func (a *AreaHandler) getAreasInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Areas get")
}

func (a *AreaHandler) getAreaInfo(w http.ResponseWriter, r *http.Request) {
	areaId := r.PathValue("id")
	fmt.Fprintln(w, "Area get "+areaId)
}
