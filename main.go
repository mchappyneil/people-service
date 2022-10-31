package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mchappyneil/people-service/server"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/users/:name", server.GetPerson)
	e.GET("/users", server.GetPeople)
	e.POST("/users", server.SavePerson)
	e.PUT("/users/:name", server.UpdatePersonAddress)
	e.DELETE("/users/:name", server.DeletePerson)
	e.Logger.Fatal(e.Start(":1324"))
}
