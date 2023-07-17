package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"name":    "123",
	// 		"address": "3",
	// 	})
	// })

	// e.GET("/about", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]string{
	// 		"message": "Halo nama saya Angga",
	// 	})
	// })

	e.Logger.Fatal(e.Start("localhost:5000"))
}
