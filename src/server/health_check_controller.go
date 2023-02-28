package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheckController はヘルスチェックコントローラーを表す。
type HealthCheckController struct {
	Version, Revision, BuiltAt, GoVersion string
}

func (c *HealthCheckController) apply(s *APIServer) {
	s.HealthCheckController = c
}

// HealthCheck はヘルスチェックする。
func (c *HealthCheckController) Healthcheck(ctx echo.Context) error {
	r := &struct {
		Version, Revision, Built, Go string
	}{
		c.Version, c.Revision, c.BuiltAt, c.GoVersion,
	}
	return ctx.JSON(http.StatusNoContent, r)
}
