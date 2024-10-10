package main

import (
	"api-with-tdd/config"
	"api-with-tdd/routes"
)

func main() {
	db := config.InitDatabase("host=localhost user=postgres password=mypassword port=4444 dbname=api_with_tdd")
	app := routes.InitRoutes(db)

	app.Run(":8000")
}
