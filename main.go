package main

import (
	mysqldb "ecommerce/connection"
	"ecommerce/routes"
	"log"
)

func main() {

	db := mysqldb.SetupDB()
	r := routes.SetupRoutes(db)
	log.Println("listening on http://localhost:5000")
	r.Run(":5000")

}
