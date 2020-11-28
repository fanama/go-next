package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	port := "3000"

	fmt.Println("starting service")

	app := fiber.New()

	app.Listen(":" + port)

}
