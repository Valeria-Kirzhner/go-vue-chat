package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
    app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID: "1512461",
		Key: "88553fa8f6ec1411c857",
		Secret: "10bea6d6ac644285766a",
		Cluster: "ap2",
		Secure: true,
	  }
	  
    app.Post("/api/messages", func (c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil{
			return err
		}
		  pusherClient.Trigger("chat", "message", data)// channel, event, data


        return c.JSON([]string{})
    })

    log.Fatal(app.Listen(":8000"))
}