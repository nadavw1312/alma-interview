package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		Errors []ErrorValidatorField `json:"errors"`
	}

	ErrorValidatorField struct {
		FailedField string      `json:"failedField"`
		Tag         string      `json:"tag"`
		Value       interface{} `json:"value"`
		Message     string      `json:"message"`
	}
)

var Errors = &struct {
	InvalidCredentials string
}{
	InvalidCredentials: "Invalid credentials",
}

func GetError(ctx *fiber.Ctx, status int, err string) error {
	return ctx.Status(status).JSON(NewErrorResponse(err))
}

func NewErrorResponse(err string) ErrorResponse {
	var elem ErrorValidatorField
	elem.Message = err
	fmt.Printf(err)

	return ErrorResponse{[]ErrorValidatorField{elem}}
}
