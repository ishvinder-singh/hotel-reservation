package main

import (
	"context"
	"flag"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ishvinder-singh/hotel-reservation/api"
	"github.com/ishvinder-singh/hotel-reservation/db"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	listenAddr := flag.String("listenAddr", ":5002", "The listen address of the API server")
	flag.Parse()
	app := fiber.New(config)
	apiv1 := app.Group("api/v1")

	//handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	apiv1.Get("/user", userHandler.HandleGetUsers)

	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)

}
