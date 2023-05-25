package handlers

import (
	"backend_template/src/core/domain/errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/wallrony/go-validator/validator"
)

var unprocessableEntityError = &echo.HTTPError{
	Code: http.StatusUnprocessableEntity,
}
var unsupportedMediaTypeError = &echo.HTTPError{
	Message: "Unsupported Media Type",
	Code:    http.StatusUnsupportedMediaType,
}
var badRequestError = &echo.HTTPError{
	Code: http.StatusBadRequest,
}
var internalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "Ocorreu um erro inesperado. Por favor, contate o suporte.",
}
var unauthorizedError = &echo.HTTPError{
	Code: http.StatusUnauthorized,
}

func badRequestErrorWithMessage(message string) *echo.HTTPError {
	err := badRequestError
	err.Message = message
	return err
}

func unprocessableEntityErrorWithMessage(message string) *echo.HTTPError {
	err := unprocessableEntityError
	err.Message = message
	return err
}

func unsupportedMediaTypeErrorWithMessage(message string) *echo.HTTPError {
	err := unsupportedMediaTypeError
	err.Message = message
	return err
}

func responseFromError(err errors.Error) error {
	var e *echo.HTTPError = badRequestError
	if err.CausedInternally() {
		e = internalServerError
	} else if err.CausedByValidation() {
		e = unprocessableEntityError
	}
	e.Message = strings.Join(err.Messages(), ";")
	return e
}

func responseFromValidationError(valErr validator.ValidationError) error {
	var e *echo.HTTPError = badRequestError
	var err = errors.NewValidation(valErr.Messages())
	if err.CausedInternally() {
		e = internalServerError
	} else if err.CausedByValidation() {
		e = unprocessableEntityError
	}
	e.Message = strings.Join(err.Messages(), ";")
	return e
}

func getUUIDParamFromRequestPath(ctx echo.Context, paramName string) (*uuid.UUID, errors.Error) {
	strUUID := ctx.Param(paramName)
	if strUUID == "" {
		return nil, errors.NewFromString(fmt.Sprintf("you must provide a valid %s", paramName))
	} else if uuid, err := uuid.Parse(strUUID); err != nil {
		return nil, errors.NewFromString("the provided id is invalid")
	} else {
		return &uuid, nil
	}
}
