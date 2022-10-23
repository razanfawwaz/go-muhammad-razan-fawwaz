package main

import (
	"weekly-task2/config"
	"weekly-task2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
