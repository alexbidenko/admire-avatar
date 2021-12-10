package main

import (
	"admire-avatar/config"
	"admire-avatar/entities"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Version is " + config.Version)
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

	go func() {
		files, err := ioutil.ReadDir("files/temporary")
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, file := range files {
			if file.Mode().IsRegular() && filepath.Ext(file.Name()) == "png" && time.Now().Sub(file.ModTime()) > 2*time.Hour {
				err = os.Remove("files/images/" + file.Name())
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	fmt.Println("Application started")
	port := "7015"
	if os.Getenv("MODE") == "production" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, routes))
}

func migrate() {
	err := config.DB.AutoMigrate(&entities.User{}, &entities.Image{})
	if err != nil {
		panic(err)
	}
}
