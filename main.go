package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/msayib/fibernote/controllers"
	"github.com/msayib/fibernote/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/notes", func(router fiber.Router) {
		router.Post("/", controllers.Store)
		router.Get("", controllers.Index)
	})
	micro.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", controllers.Destroy)
		router.Get("", controllers.Show)
		router.Put("", controllers.Update)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Fibernote | @ibb.ac | github/msayib",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
