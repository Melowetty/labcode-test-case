package dto

import (
	"github.com/ansel1/merry"
	"net/http"
)

var (
	BaseError            = merry.New("Base error")
	NotFoundError        = BaseError.WithHTTPCode(http.StatusNotFound).WithMessage("Entity not found")
	AreaNotFoundError    = NotFoundError.WithUserMessage("Area not found")
	CameraNotFoundError  = NotFoundError.WithUserMessage("Camera not found")
	CameraNotInAreaError = BaseError.WithHTTPCode(http.StatusInternalServerError).WithMessage("Camera not in area").WithUserMessage("Camera not in area")
)
