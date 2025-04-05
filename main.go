package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"labcode-test-case/internal/handler"
	"labcode-test-case/internal/service"
	"labcode-test-case/internal/storage"
	"net/http"
	"os"
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
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	mux := http.NewServeMux()
	validate := validator.New()

	cameraStorage := storage.NewCameraStorage(pool)
	cameraService := service.NewCameraService(cameraStorage)

	areaStorage := storage.NewAreaStorage(pool, cameraStorage)
	areaService := service.NewAreaService(areaStorage)

	server := &Server{}
	server.cameraHandler = handler.NewCameraHandler(mux, validate, cameraService)
	server.areaHandler = handler.NewAreaHandler(mux, validate, areaService)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
