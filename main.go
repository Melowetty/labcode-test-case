package main

import (
	"fmt"
	"labcode-test-case/internal/handler"
	"net/http"
)

type Server struct {
	cameraHandler *handler.CameraHandler
	areaHandler   *handler.AreaHandler
}

func main() {
	mux := http.NewServeMux()

	server := &Server{}
	server.cameraHandler = handler.NewCameraHandler(mux)
	server.areaHandler = handler.NewAreaHandler(mux)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
