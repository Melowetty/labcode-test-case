package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/handler"
	"labcode-test-case/internal/service"
	"net/http"
)

type Server struct {
	cameraHandler *handler.CameraHandler
	areaHandler   *handler.AreaHandler
}

// @title Labcode test case
// @version 1.0
// @description This is solution of labcode test case

// @contact.name Melowetty
// @contact.url https://github.com/melowetty
// @contact.email melowetty@mail.ru

// @host localhost:8080
// @BasePath /
func main() {
	mux := http.NewServeMux()
	validate := validator.New()

	areaService := service.AreaService{}

	server := &Server{}
	server.cameraHandler = handler.NewCameraHandler(mux, validate)
	server.areaHandler = handler.NewAreaHandler(mux, validate, &areaService)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
