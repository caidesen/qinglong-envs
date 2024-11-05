package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("qinglong-envs")
	})
	log.Fatal(app.Listen("0.0.0.0:3000"))
}
