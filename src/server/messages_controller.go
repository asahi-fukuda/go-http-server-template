package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myuoncorp/go-http-server-template/logger"
	"github.com/myuoncorp/go-http-server-template/oapistub"
	"github.com/myuoncorp/go-http-server-template/server/adapter"
	"github.com/myuoncorp/go-http-server-template/usecase/message"
)

type MessagesController struct {
	ListMessagesUseCase *message.List
	SaveMessagesUseCase *message.Save
}

func (c *MessagesController) apply(s *APIServer) {
	s.MessagesController = c
}

func (s *MessagesController) GetMessages(ctx echo.Context) error {
	output, err := s.ListMessagesUseCase.Execute(ctx.Request().Context(), &message.ListInput{})
	if err != nil {
		logger.Errorf("failed to list messages. err: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "unexpected error occurred.",
		})
	}

	var response oapistub.GetMessagesSuccess
	response.Messages = adapter.ToOapiMessage(output.Messages)
	return ctx.JSON(http.StatusOK, response)
}

func (s *MessagesController) CreateMessage(ctx echo.Context) error {
	req := &oapistub.NewMessage{}
	err := ctx.Bind(req)
	if err != nil {
		logger.Errorf("failed to get message. err: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "unexpected error occurred.",
		})
	}

	output, err := s.SaveMessagesUseCase.Execute(ctx.Request().Context(), &message.SaveInput{
		Name:    req.Name,
		Message: req.Message,
	})
	if err != nil {
		logger.Errorf("failed to get message. err: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "unexpected error occurred.",
		})
	}

	return ctx.JSON(http.StatusOK, oapistub.CreateMessagesSuccess{Id: output.ID})
}
