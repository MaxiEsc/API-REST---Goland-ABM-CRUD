package main

import (
	"goapirest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Gracias a la nueva libreria de GorilaMux se utiliza de esta manera se crea un mux dedidacado a contener 
los los handlers en donde permiten realizar los endpoint con elmetodo hadleFunc  y con los metodods finales de 
"GET", "POST", "PUT", "DELETE"


*/

func main() {
	//Rutas para MUX
	mux := mux.NewRouter()

	//EndPoint
	mux.HandleFunc("/api/user/", handlers.GetUsuarios).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUsuario).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUsuario).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUsuario).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUsuario).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
