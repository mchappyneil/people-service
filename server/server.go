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
	return c.JSON(http.StatusNotFound, "Person not found")
}

func savePerson(c echo.Context) error {
	p := new(Person)
	if err := c.Bind(&p); err != nil {
		return err
	}
	people = append(people, *p)
	return c.JSON(http.StatusCreated, p)
}

//func updatePersonName(c echo.Context) error {
//	var personPointer *Person
//	var newPerson Person
//	var index int
//	name := c.Param("name")
//	for i, p := range people {
//		if p.Name == name {
//			personPointer = &p
//			index = i
//		}
//	}
//	if personPointer == nil {
//		return c.JSON(http.StatusNotFound, "Person not found")
//	}
//	people = splice(people, index)
//	if err := c.Bind(&newPerson); err != nil {
//		return err
//	}
//	people = append(people, *newPerson)
//	(*personPointer).Name = newPerson.Name
//	return c.JSON(http.StatusOK, *personPointer)
//}
//
func deletePerson(c echo.Context) error {
	var person *Person
	var index int
	name := c.Param("name")
	for i, p := range people {
		if p.Name == name {
			person = &p
			index = i
		}
	}
	if person == nil {
		return c.JSON(http.StatusNotFound, "Person not found")
	}
	people = splice(people, index)
	return c.JSON(http.StatusOK, person)
}

func splice(slice []Person, index int) []Person {
	return append(slice[:index], slice[index+1:]...)
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
	e.DELETE("/users/:name", deletePerson)
	e.Logger.Fatal(e.Start(":1323"))
}
