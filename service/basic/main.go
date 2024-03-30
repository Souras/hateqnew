package main

import (
	"database/sql" // add this
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // add this

	"github.com/gofiber/fiber/v2"
)

func main() {
	// connStr := "postgresql://casaos:casaos@192.168.1.12/todos?sslmode=disable"
	// // Connect to database
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var err error
	db, err := sql.Open("postgres", "user=casaos dbname=todos password=casaos host=192.168.1.12 port=5432 sslmode=disable") // Replace credentials
	if err != nil {
		fmt.Println("In Db error")
		panic(err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) string[] {
		return indexHandler1(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func indexHandler1(c *fiber.Ctx, db *sql.DB) []string {
	var res string
	var todos []string
	fmt.Println("In indexHandler")
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	fmt.Println(todos)
	return todos
}
