package main

import (
	"./bolter"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {

	e := echo.New()
	e.Get("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Propr")
	})

	e.Post("/create/:appname", func(c echo.Context) error {

		appname := c.Param("appname")

		bolter.Create(appname)

		return c.HTML(http.StatusOK, "Done")
	})

	e.Post("/add/:appname/:key/:value", func(c echo.Context) error {

		appname := c.Param("appname")
		key := c.Param("key")
		value := c.Param("value")

		bolter.Put(appname, key, value)

		return c.HTML(http.StatusOK, "Done")
	})

	e.Get("/get/:appname/:key", func(c echo.Context) error {

		appname := c.Param("appname")
		key := c.Param("key")

		value := bolter.Get(appname, key)

		return c.HTML(http.StatusOK, value)
	})

	e.Get("/get/:appname", func(c echo.Context) error {

		appname := c.Param("appname")

		value := bolter.GetAll(appname)

		return c.HTML(http.StatusOK, value)
	})

	e.Run(standard.New(":1323"))
}
