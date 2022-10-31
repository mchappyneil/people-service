package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mchappyneil/people-service/models"
	"net/http"
	"net/url"
)

// const PersonNotFound = "Error - person not found: "

var people = []models.Person{
	{Name: "John Doe", Address: "123 Foo Street"},
	{Name: "Neil Mehta", Address: "321 Bar Avenue"},
}

type Links struct {
	Self string `json:"self"`
}

type JSONAPIResponse struct {
	Links `json:"links"`
	Data  interface{} `json:"data"`
}

//func createError(errorString string, name string) error {
//	return fmt.Errorf("%s%q", errorString, name)
//}

func GetPeople(c echo.Context) error {
	return c.JSON(http.StatusOK, people)
}

func GetPerson(c echo.Context) error {
	name := c.Param("name")
	for _, p := range people {
		if p.Name == name {
			t := &url.URL{Path: p.Name}
			encodedName := t.String()
			response := JSONAPIResponse{
				Data: p,
				Links: Links{
					//Figure out how to escape for a URL as in space becomes %20 - DONE
					Self: fmt.Sprintf("localhost:1324/%s", encodedName),
				},
			}
			return c.JSON(http.StatusOK, response)
		}
	}
	return c.JSON(http.StatusNotFound, "Person not found")
}

func SavePerson(c echo.Context) error {
	p := new(models.Person)
	c.Echo().Validator = &models.PersonValidator{Validator: models.V}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Validate(p); err != nil {
		return err
	}
	people = append(people, *p)
	return c.JSON(http.StatusCreated, p)
}

func UpdatePersonAddress(c echo.Context) error {
	var newPerson models.Person
	c.Echo().Validator = &models.PersonValidator{Validator: models.V}
	if err := c.Bind(&newPerson); err != nil {
		return err
	}
	// Validate data
	// Validate name not empty
	if err := c.Validate(newPerson); err != nil {
		return err
	}

	// Pass to PersonService

	// Manipulate data into JSON API response

	// Technically GetUser
	index := -1
	name := c.Param("name")
	for i, p := range people {
		if p.Name == name {
			index = i
			break
		}
	}
	if index < 0 {
		return c.JSON(http.StatusNotFound, "Person not found")
	}
	////

	// Technically UpdateUser
	people[index] = newPerson

	////
	return c.JSON(http.StatusOK, JSONAPIResponse{Data: newPerson})
}

func DeletePerson(c echo.Context) error {
	var person *models.Person
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

func splice(slice []models.Person, index int) []models.Person {
	return append(slice[:index], slice[index+1:]...)
}
