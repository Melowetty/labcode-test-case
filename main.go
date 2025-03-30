package main

import (
	"fmt"
	"labcode-test-case/internal/controller"
	"net/http"
)

type Server struct {
	cameraController *controller.CameraController
	areaController   *controller.AreaController
}

func main() {
	server := &Server{}
	server.cameraController = controller.NewCameraController()
	server.areaController = controller.NewAreaController()

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
