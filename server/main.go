package main

import (
	"admire-avatar/config"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("MODE") != "" {
		fmt.Println("Initializing with mode " + os.Getenv("MODE"))
	} else {
		fmt.Println("Initializing")
	}
	err := config.DBError
	if err != nil {
		panic(err)
	}
	migrate()

	fmt.Println("Session created")

	routes := initRoutes()

	http.Handle("/", routes)
	fmt.Println("Application started")
	port := "7015"
	if os.Getenv("MODE") == "production" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, routes))
}

func migrate() {

}
