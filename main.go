package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


var DB *gorm.DB
func main() {
	// Create instance of Fiber
	app := fiber.New()

	// Create HTTP handler
	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "go fiber first created api",
		})
	})

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	godotenv.Load()
	dbhost := os.Getenv(key:"MYSQL_HOST")
	dbuser := OS.Getenv(key:"MYSQL_USER")
	dbpassword := os.Getenv(key:"MYSQL_PASSWORD")
	dbname := os.Getenv(key:"MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil{
		panic("db connection failed")
	}

	DB = db
	fmt.Println("db connected successfully")
}
