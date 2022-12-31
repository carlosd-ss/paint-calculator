package server

import (
	"digitalrepublic/api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"os"
	"os/signal"
)

type Server interface {
	Start()
}

type server struct {
	Fiber *fiber.App
}

func New() Server {
	return &server{}
}

func (e *server) Start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = e.Fiber.Shutdown()
	}()

	e.Fiber = fiber.New(fiber.Config{
		AppName:      "Paint Calculator!",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	e.Fiber.Use(cors.New())
	e.Fiber.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))

	e.Fiber.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Paint Calculator!"))
	})
	api := e.Fiber.Group("/api/v1")
	routes.Router(api)

	// Prepare an endpoint for 'Not Found'.
	e.Fiber.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 8080.
	e.Fiber.Listen(":8080")
}
