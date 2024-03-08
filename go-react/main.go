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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "*",
	}))

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

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")

		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID de usuario no válido",
			})
		}

		filter := bson.M{"_id": objID}

		coll := client.Database("go-react").Collection("users")

		_, err = coll.DeleteOne(context.TODO(), filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al eliminar usuario",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Usuario eliminado exitosamente",
		})
	})

	// Handler para editar un usuario
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")
		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID de usuario no válido",
			})
		}
		var updatedUser models.User
		if err := c.BodyParser(&updatedUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error al analizar los datos del usuario",
			})
		}

		coll := client.Database("go-react").Collection("users")

		filter := bson.M{"_id": objID}
		update := bson.M{"$set": bson.M{
			"Name": updatedUser.Name,
		}}
		ee, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al actualizar el usuario",
			})
		}
		fmt.Printf("UpdateResult: %+v\n", ee)

		return c.JSON(fiber.Map{
			"message": "Usuario actualizado exitosamente",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}
