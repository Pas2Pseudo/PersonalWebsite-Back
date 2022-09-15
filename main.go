package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var config = GetConfig()

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	var redisConfig = config.RedisConfig
	redisManager := newRedisManager(redisConfig.Address, redisConfig.Port, redisConfig.Password, redisConfig.DB)

	redisManager.set("test", "Hello World !")
	println(redisManager.get("test"))

	err := app.Listen(":" + strconv.Itoa(config.Port))
	if err != nil {
		return
	}
}