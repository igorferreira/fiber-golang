package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)


type healthStatus struct {
    Status   string `json:"status"`
}


func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 5")
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error  {
		c.Set("Connection","keep-alive")
		c.Set("Content-Type","text/html; charset=utf-8")
		resp := " Welcome to fiber, dear " + c.Params("name") + "!"
		return c.SendString(resp)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		c.Set("Content-type","application/json")
		status := &healthStatus{ Status:   "OK"}
		statusJSON, _ := json.Marshal(status)
		fmt.Println(string(statusJSON))
		return c.SendString(string(statusJSON))
	})

	app.Static("/app","./public")

	app.Listen(":8080")
}