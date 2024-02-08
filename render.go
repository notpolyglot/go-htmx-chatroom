package main

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")

	return component.Render(context.Background(), c.Response())

}
