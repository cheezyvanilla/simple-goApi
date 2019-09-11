package main

import(
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.Use(tokenAuth) //middleware for user's token check
	router.HandleFunc("/signup", signUp).Methods("POST")
	router.HandleFunc("/signin", signIn)
	router.HandleFunc("/tweet", tweet)
	log.Fatal(http.ListenAndServe(":8080", router))

	

}