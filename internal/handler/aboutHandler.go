package handler

import (
	"github.com/joshbrgs/my-portfolio/internal/template/about"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct{}

func (h *AboutHandler) HandleAbout(c echo.Context) error {
	return about.About().Render(c.Request().Context(), c.Response())
}
