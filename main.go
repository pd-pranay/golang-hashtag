package main

import (
	"hashtag/controllers"
	"hashtag/db"
	"hashtag/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	dbEngine := "postgres"
	dbUser := "hashtag"
	dbPassword := "hashtag"
	dbName := "hashtag"
	dbHost := "localhost"
	dbPort := "7777"
	dbSSLMode := "disable"

	// Connect to Database
	dbCon, err := db.Connect(dbEngine, dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()

	queries := db.New(dbCon)

	postController := controllers.NewPostsController(dbCon, queries)

	app := fiber.New()

	v1 := app.Group("/v1/api")

	routes.PostsRoutes(v1, postController)

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")

}
