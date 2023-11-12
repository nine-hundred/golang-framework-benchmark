package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Item struct {
	Id   int    `json:"id" query:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/item", Get)
	e.POST("/item", Post)
	e.Logger.Fatal(e.Start(":8080"))
}

func Get(c echo.Context) error {
	var item Item
	echoBinder := &echo.DefaultBinder{}
	err := echoBinder.BindQueryParams(c, &item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	item.Name = "hello"
	return c.JSON(http.StatusOK, item)
}

func Post(c echo.Context) error {
	echoBinder := &echo.DefaultBinder{}
	item := new(Item)

	err := echoBinder.BindBody(c, item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, item)
}
