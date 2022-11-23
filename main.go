package main

import (
	"capstone-project/config"
	"capstone-project/entity"
	"capstone-project/repository"
	"capstone-project/route"
	"capstone-project/service"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		db         = config.InitDatabase()
		PORT       = os.Getenv("PORT")
		repository = repository.NewRepository(db)
		service    = service.NewService(repository)
	)
	entity.DB.AutoMigrate()
	app := route.New(service)

	app.Start(":" + PORT)
}
