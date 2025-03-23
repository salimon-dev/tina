package main

import (
	"net/http"
	"salimon/tina/middlewares"
	"salimon/tina/types"

	"github.com/labstack/echo/v4"
)

func InteractHandler(ctx echo.Context) error {
	payload := new(types.InteractSchema)
	if err := ctx.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	// validation errors
	vError, err := middlewares.ValidatePayload(*payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	if vError != nil {
		return ctx.JSON(http.StatusBadRequest, vError)
	}

	response := types.Message{
		From: "tina",
		Type: types.MessageTypePlain,
		Body: "hello!",
	}
	return ctx.JSON(http.StatusOK, response)
}
