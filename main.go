package main

import (
	"fmt"

	"github.com/amirdaraby/golang-mvc/config"
	"github.com/amirdaraby/golang-mvc/database"
	"github.com/amirdaraby/golang-mvc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	cfg, err := config.Init()

	if err != nil {
		panic(err)
	}

	err = database.Init(&cfg.Database)

	db := database.DbConnection

	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()

		if err != nil {
			fmt.Println(err)
		}

		err = sqlDB.Close()

		if err != nil {
			fmt.Println(err)
		}
	}()

	app := fiber.New(fiber.Config{})

	routes.Register(app)

	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} | ${ip} \n",
	}))

	panic(app.Listen(":8000"))
}
