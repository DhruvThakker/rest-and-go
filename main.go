//Taken from 
//https://medium.com/@gautamprajapati/writing-a-simple-e-commerce-website-api-in-go-programming-language-9f671324b4dd
//Entry point for API
package main

import(
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"rest-and-go/store"
)

func main(){
	//Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := store.NewRouter() //create routes

	// These two lines are import for cross origin
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"})

	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins,allowedMethods)(router)))
}