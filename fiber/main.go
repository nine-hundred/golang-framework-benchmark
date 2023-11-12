package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := fiber.New()
	app.Get("/item", Get)
	app.Post("/item", Post)
	app.Listen(":8080")
}

func Get(c *fiber.Ctx) error {
	var item Item
	item.Id = c.QueryInt("id")
	item.Name = "hello"

	return c.JSON(item)
}

func Post(c *fiber.Ctx) error {
	item := new(Item)

	if err := json.Unmarshal(c.BodyRaw(), item); err != nil {
		return err
	}

	return c.JSON(item)
}
