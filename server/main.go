package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("user", func(c *fiber.Ctx) error {

		val := ""

		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)

		millisec := random.Intn(4000)

		if millisec > 2000 {
			val = fmt.Sprintf("Time: %v | [exceeded]", millisec)
		} else {
			val = fmt.Sprintf("Time: %v", millisec)
		}

		time.Sleep(time.Millisecond * time.Duration(millisec))

		fmt.Println(val)

		return c.SendString(val)

	})

	app.Get("products", func(c *fiber.Ctx) error {

		val := ""

		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)

		millisec := random.Intn(4000)

		if millisec > 2000 {
			val = fmt.Sprintf("Time: %v | [exceeded]", millisec)
		} else {
			val = fmt.Sprintf("Time: %v", millisec)
		}

		time.Sleep(time.Millisecond * time.Duration(millisec))

		fmt.Println(val)

		return c.SendString(val)

	})

	app.Listen(":5000")

}
