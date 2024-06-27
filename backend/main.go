package main

import (
	"kenshilabs_demo/controllers"
	"kenshilabs_demo/database"
	"kenshilabs_demo/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	database.InitConnection()
}
func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, token, Authorization",
		AllowMethods: "GET, POST, PATCH, DELETE, OPTIONS, PUT",

		AllowCredentials: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent) // 204 No Content
		}
		return c.Next()
	})

	app.Route("/", func(router fiber.Router) {
		router.Post("/signup", controllers.Signup)
		router.Post("/signin", controllers.SignIn)
		router.Post("/signout", controllers.SignOut)
	})

	task := fiber.New()
	app.Mount("/api", task)

	task.Route("/tasks", func(router fiber.Router) {
		router.Post("/", middleware.AuthMiddleware, controllers.CreateTask)
		router.Get("", controllers.GetTask)
	})
	task.Route("/tasks/:id", func(router fiber.Router) {
		router.Delete("", middleware.AuthMiddleware, controllers.DeleteTask)
		router.Get("", middleware.AuthMiddleware, controllers.GetTaskById)
		router.Put("", middleware.AuthMiddleware, controllers.UpdateTask)
	})

	log.Fatal(app.Listen(":8000"))
}
