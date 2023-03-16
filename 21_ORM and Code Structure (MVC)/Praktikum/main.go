package main

import (
	"rest/config"
	"rest/routes"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	config.InitialMigration()

	e := routes.New()
	e.Logger.Fatal(e.Start(":9000"))
}
