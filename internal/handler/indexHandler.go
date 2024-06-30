package handler

import (
	"github.com/joshbrgs/my-portfolio/internal/template/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func (h *IndexHandler) HandleIndex(c echo.Context) error {
	return index.Index().Render(c.Request().Context(), c.Response())
}
