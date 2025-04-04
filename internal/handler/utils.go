package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"labcode-test-case/internal/handler/model"
	"net/http"
	"strconv"
	"strings"
)

const (
	InternalServerErrorMessage = "Internal server error"
)

func parsePathValueAsInt(validate *validator.Validate, r *http.Request, path string) (int, error) {
	value, err := strconv.Atoi(r.PathValue(path))

	if err != nil {
		return 0, err
	}

	validationErr := validate.Var(value, "required,gt=0")

	if validationErr != nil {
		return 0, validationErr
	}

	return value, nil
}

func parseBody(r *http.Request, target any) error {
	err := json.NewDecoder(r.Body).Decode(&target)

	if err != nil {
		return err
	}

	return nil
}

func validateBody(validate *validator.Validate, body interface{}) validator.ValidationErrors {
	err := validate.Struct(body)

	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		return validationErrors
	}

	return nil
}

func writeValidationErrorsResponse(w http.ResponseWriter, validationErrs validator.ValidationErrors) {
	var failedFields []string
	for _, validationError := range validationErrs {
		failedFields = append(failedFields, validationError.StructField())
	}
	fmt.Printf(validationErrs.Error())
	message := fmt.Sprintf("Validation failed for fields: %s", strings.Join(failedFields[:], ", "))
	writeError(w, message, http.StatusBadRequest)
}

func writeError(w http.ResponseWriter, message string, status int) {
	customErr := &model.CustomError{Message: message}
	writeJsonResponse(w, customErr, status)
}

func writeOkJsonResponse(w http.ResponseWriter, data interface{}) {
	writeJsonResponse(w, data, 200)
}

func writeJsonResponse(w http.ResponseWriter, data interface{}, status int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

func writeInternalServerError(w http.ResponseWriter) {
	writeError(w, InternalServerErrorMessage, http.StatusInternalServerError)
}
