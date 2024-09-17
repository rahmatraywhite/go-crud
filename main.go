package main

import (
	"go-web-native/config"
	categoryController "go-web-native/controllers/category"
	homeController "go-web-native/controllers/home"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/", homeController.Welcome)

	// Categori Handler
	http.HandleFunc("/categories", categoryController.Index)
	http.HandleFunc("/categories/add", categoryController.Add)
	http.HandleFunc("/categories/edit", categoryController.Edit)
	http.HandleFunc("/categories/delete", categoryController.Delete)

	log.Println("Server running on port 4000")
	http.ListenAndServe(":4000", nil)
}
