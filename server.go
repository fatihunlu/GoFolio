package main

import (
	"gofolio/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	root := e.Group("")
	api.InitRoutes(root)

	e.Logger.Fatal(e.Start(":1323"))
}
