package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var people = []Person{
	{Name: "John Doe", Address: "123 Foo Street"},
	{Name: "Neil Mehta", Address: "321 Bar Avenue"}}

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func getPeople(c echo.Context) error {
	return c.JSON(http.StatusOK, people)
}

func getPerson(c echo.Context) error {
	name := c.Param("name")
	for _, p := range people {
		if p.Name == name {
			return c.JSON(http.StatusOK, p)
		}
	}
	return echo.ErrBadRequest
}

func savePerson(c echo.Context) error {
	p := new(Person)
	if err := c.Bind(p); err != nil {
		return err
	}
	people = append(people, *p)
	return c.JSON(http.StatusCreated, p)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/users/:name", getPerson)
	e.GET("/users", getPeople)
	e.POST("/users", savePerson)
	//e.PUT("/users/:name", updatePersonName)
	//e.DELETE("/users/:name", deletePerson)
	e.Logger.Fatal(e.Start(":1323"))
}
