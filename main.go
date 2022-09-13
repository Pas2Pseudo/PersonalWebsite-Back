package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})


	redisManager := newRedisManager("127.0.0.1", 6379, "", 0)

	redisManager.set("test", "Hello World !")
	println(redisManager.get("test"))

	app.Listen(":8393")
}