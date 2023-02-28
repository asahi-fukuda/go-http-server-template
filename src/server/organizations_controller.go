package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/myuoncorp/go-http-server-template/logger"
	"github.com/myuoncorp/go-http-server-template/oapistub"
	"github.com/myuoncorp/go-http-server-template/server/adapter"
	"github.com/myuoncorp/go-http-server-template/usecase/organization"
)

// OrganizationsController は Organization のコントローラーを表す。
type OrganizationsController struct {
	ListOrganizationsUseCase *organization.List
	GetOrganizationsUseCase  *organization.Get
}

func (c *OrganizationsController) apply(s *APIServer) {
	s.OrganizationsController = c
}

// ListOrganizations は Organization を一覧する。
func (s *OrganizationsController) ListOrganizations(ctx echo.Context) error {
	res, err := s.ListOrganizationsUseCase.Execute(ctx.Request().Context(), &organization.ListInput{})
	if err != nil {
		logger.Errorf("failed to list organizations. err: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "unexpected error occurred.",
		})
	}

	var response oapistub.ListOrganizationsSuccess
	response.Organizations = adapter.OrganizationList(res.Organizations)
	return ctx.JSON(http.StatusOK, response)
}

// GetOrganizations は Organization を取得する。
func (s *OrganizationsController) GetOrganization(ctx echo.Context, id oapistub.OrganizationId) error {
	res, err := s.GetOrganizationsUseCase.Execute(ctx.Request().Context(), &organization.GetInput{})
	if err != nil {
		logger.Errorf("failed to get organization. err: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "unexpected error occurred.",
		})
	}

	var response oapistub.GetOrganizationSuccess
	response.Organization = adapter.Organization(res.Organization)
	return ctx.JSON(http.StatusOK, response)
}
