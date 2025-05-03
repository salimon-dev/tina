package main

import (
	"fmt"
	"io"
	"net/http"
	"salimon/tina/helpers"
	"salimon/tina/openai"

	"github.com/labstack/echo/v4"
	"github.com/salimon-dev/gomsg"
)

func InteractHandler(ctx echo.Context) error {
	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}

	schema, errs := gomsg.ParseInteractionSchema(requestBody)

	if errs != nil {
		return ctx.JSON(http.StatusBadRequest, errs)
	}

	data, err := openai.SendCompletionRequest(schema)
	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}

	result := gomsg.InteractionSchema{
		Data: []gomsg.Message{*data},
	}
	return ctx.JSON(http.StatusOK, result)
}
