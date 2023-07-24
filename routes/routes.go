package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	baseUrl = "/api/personalities"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentType)

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc(baseUrl+"", controllers.AllPersonalities).Methods("Get")
	router.HandleFunc(baseUrl+"/{id}", controllers.GetPersonality).Methods("Get")
	router.HandleFunc(baseUrl+"", controllers.CreatePersonality).Methods("Post")
	router.HandleFunc(baseUrl+"/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc(baseUrl+"/{id}", controllers.UpdatePersonality).Methods("Put")

	log.Fatal(
		http.ListenAndServe(
			":8000",
			handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router),
		),
	)

}
