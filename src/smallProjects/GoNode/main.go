package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {

	app := fiber.New()

	todos := []Todo{}


	app.Get("/api/todos", func(c *fiber.Ctx)  error {
		return c.Status(http.StatusOK ).JSON(todos)
	})

app.Post("/api/todos", func (c * fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
	}

	todo.ID = len(todos) + 1
	todos = append(todos, *todo)

	// var x int = 5 	// 0x0001 // adres w pamieci
	// var p *int = &x // 0x0001 // adres w pamieci zmiennej x

	// fmt.Println(p) // 0x0001 // adres w pamieci zmiennej x
	// fmt.Println(*p) // 5	 // wartość zmiennej x


	return c.Status(http.StatusCreated).JSON(todo)
})


app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = true
			return c.Status(http.StatusOK).JSON(todos[i])
		}
	}
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
})

app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i + 1:]...)
			return c.Status(http.StatusOK).JSON(todos)
		}
	}
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
})

	log.Fatal(app.Listen(":3000"))
}