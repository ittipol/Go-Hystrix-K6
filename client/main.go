package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func init() {
	hystrix.ConfigureCommand("user", hystrix.CommandConfig{
		Timeout:                1500,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  80,
		SleepWindow:            10000,
	})

	hystrix.ConfigureCommand("products", hystrix.CommandConfig{
		Timeout:                1500,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  80,
		SleepWindow:            10000,
	})

	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()
	go http.ListenAndServe(":8081", streamHandler)
}

func main() {

	app := fiber.New()

	app.Get("user", getUser)

	app.Get("products", getProducts)

	app.Listen(":8080")

}

func getUser(c *fiber.Ctx) error {

	fmt.Println("[getUser] Requesting....")

	output := make(chan bool, 1)

	hystrix.Go("user", func() error {

		res, err := httpGet("http://server:5000/user")

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return fiber.ErrInternalServerError
		}

		fmt.Printf("Res: %v\n", res)

		output <- true

		return nil

	}, func(err error) error {

		fmt.Printf("Error: %v\n", err)

		return nil
	})

	select {
	case v := <-output:

		if !v {
			return fiber.ErrInternalServerError
		}

	}

	c.Status(fiber.StatusOK)
	return c.JSON("OK")
}

func getProducts(c *fiber.Ctx) error {

	fmt.Println("[getProducts] Requesting....")

	output := make(chan bool, 1)

	hystrix.Go("products", func() error {

		res, err := httpGet("http://server:5000/products")

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return fiber.ErrInternalServerError
		}

		fmt.Printf("Res: %v\n", res)

		output <- true

		return nil

	}, func(err error) error {

		fmt.Printf("Error: %v\n", err)

		return nil
	})

	select {
	case v := <-output:

		if !v {
			return fiber.ErrInternalServerError
		}

	}

	c.Status(fiber.StatusOK)
	return c.JSON("OK")
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("- #1")
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("- #2")
		return "", err
	}

	return string(body), err
}
