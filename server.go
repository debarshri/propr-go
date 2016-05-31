package main

import (
	"./bolter"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"io/ioutil"
)

func main() {

	e := echo.New()
	b := bolter.New();

	e.Get("/", func(c echo.Context) error {

		b, err := ioutil.ReadFile("propr-ui/app/index.html")
		if err != nil {
			panic(err)
		}

		s := string(b[:])

		return c.HTML(http.StatusOK, s)
	})

	e.Get("/app.js", func(c echo.Context) error {

		return c.File("propr-ui/app/app.js")
	})

	e.Post("/create/:appname", func(c echo.Context) error {

		app_name := c.Param("appname")

		b.Create(app_name)

		return c.HTML(http.StatusOK, "Done")
	})

	e.Post("/add/:appname/:key/:value", func(c echo.Context) error {

		app_name := c.Param("appname")
		key := c.Param("key")
		value := c.Param("value")

		b.Put(app_name, key, value)

		return c.HTML(http.StatusOK, "Done")
	})

	e.Get("/get/:appname/:key", func(c echo.Context) error {

		app_name := c.Param("appname")
		key := c.Param("key")

		value := b.Get(app_name, key)

		return c.HTML(http.StatusOK, value)
	})


	e.Get("/get/:appname", func(c echo.Context) error {

		app_name := c.Param("appname")

		value := b.GetAll(app_name)

		return c.HTML(http.StatusOK, value)
	})

	//TODO app ui

	e.Run(standard.New(":1323"))
}
