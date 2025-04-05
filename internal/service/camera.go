package service

type CameraStorage interface {
}

type CameraService struct {
	cameraStorage CameraStorage
}

func NewCameraService(storage CameraStorage) *CameraService {
	service := &CameraService{cameraStorage: storage}
	return service
}
