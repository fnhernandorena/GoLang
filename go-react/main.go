package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fnhernandorena/go-react/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error cargando archivo .env: %v", err)
	}
	mongodb := os.Getenv("MONGODB_URI")

	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb))
	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		coll := client.Database("go-react").Collection("users")
		results, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			println(err)
		}
		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}
		return c.JSON(&fiber.Map{
			"users": users,
		})

	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)
		coll := client.Database("go-react").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key: "Name", Value: user.Name,
		}})

		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}
