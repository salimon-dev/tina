package main

import (
	"fmt"
	"net/http"
	"salimon/tina/helpers"
	"salimon/tina/middlewares"
	"salimon/tina/openai"
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

	response, err := openai.SendCompletionRequest(payload.Data)
	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}
	return ctx.JSON(http.StatusOK, response)
}
