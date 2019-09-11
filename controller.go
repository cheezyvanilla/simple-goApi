package main
import (
	//"fmt"
	"net/http"
	
)

func signUp(w http.ResponseWriter, r *http.Request){
	// post data process
	r.ParseForm()
	email := r.FormValue("email")
	pswd := r.FormValue("pswd")

	//db open
	db := connect()
	defer db.Close()

	//db process
	db.QueryRow("INSERT INTO users(email, password) VALUES($1,$2)", email, pswd)
	
}