package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

func main() {
	app := fiber.New()
	app.Use(cors.New())

	setupDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/quizzes", getQuizzes)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	app.Listen(":5000")
}

func setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	quizCollection = client.Database("quiz").Collection("quizzes")
}

func getQuizzes(c *fiber.Ctx) error {
	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	var quizzes []map[string]any

	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		return err
	}

	return c.JSON(quizzes)
}
