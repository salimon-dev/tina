package webhook

import (
	"fmt"
	"net/http"
	"salimon/tina/middlewares"
	"salimon/tina/nexus"
	"salimon/tina/openai"
	"salimon/tina/types"

	"github.com/labstack/echo/v4"
)

func Handle(ctx echo.Context) error {
	type schema struct {
		Messages []types.Message `json:"messages" validate:"required"`
	}
	payload := new(schema)
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

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	if len(payload.Messages) == 0 {
		return ctx.JSON(http.StatusOK, Response{
			Success: false,
			Message: "no messages sent",
		})
	}
	threadId := payload.Messages[0].ThreadId

	messages, err := nexus.GetLastMessages(threadId)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusOK, Response{
			Success: false,
			Message: err.Error(),
		})
	}

	result, err := openai.SendCompletionRequest(messages)
	if err != nil {
		return ctx.JSON(http.StatusOK, Response{
			Success: false,
			Message: "failed to fulfill the request",
		})
	}

	err = nexus.SendPlainMessage(threadId, result.Body)
	if err != nil {
		return ctx.JSON(http.StatusOK, Response{
			Success: false,
			Message: "failed to connect to nexus",
		})
	}

	return ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "ok",
	})
}
