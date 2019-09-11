package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.Use(tokenAuth) //middleware for user's token check
	router.HandleFunc("/signup", signUp).Methods("POST")
	router.HandleFunc("/signin", signIn).Methods("POST")
	router.HandleFunc("/tweet", tweet).Methods("POST")
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	

}