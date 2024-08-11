package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Todo struct {
	// dzieki omitempty podczas tworzenia TODO gdy Id nie jest podane, omija ta wartosc i nie dodaje jej wogóle
	//  W tym przypadku mongodb stworzy wlasne id automatycznie 
	ID 				primitive.ObjectID 		`json:"id, omitempty" bison:"id, omitempty"` 				
	Completed bool 		`json:"completed" bison:"completed"` 	
	Body 			string 	`json:"body" bison:"body"`				
}


func main() {

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There is no .env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("There is no mongodb connection", err)
	}

	collection = client.Database("golang_db").Collection("Cluster0")
	app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", createTodo)
	// app.Patch("/api/todos/:id", updateTodo)
	// app.Delete("/api/todos/:id", deleteTodo)




	fmt.Println("Connected to mongoDB")
	log.Fatal(app.Listen(os.Getenv("PORT")))
}

func getTodos(c *fiber.Ctx) error{
		var todos []Todo

		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			return err
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo Todo
			if err := cursor.Decode(&todo); err != nil {
				return err
		}

		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "The body can not be empty"})
	}
	
	insertResult, err := collection.InsertOne(context.Background(),  todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(http.StatusCreated).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Id not found"})	
	}

	filter := bson.M{"id": objId}
	update := bson.M{"$set": bson.M{"completed": "true"}}
	_, err = collection.UpdateOne(context.Background(),filter, update)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "Todo has been updated"})
}

func deleteTodo(c *fiber.Ctx) error{
	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Id not found"})	
	}

	filter := bson.M{"id": objId}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "Todo has been deleted"})
}