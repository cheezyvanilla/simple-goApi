package main
import (
	"fmt"
	"time"
	"net/http"
	"log"
	"encoding/json"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
)


//sign up handler
func signUp(w http.ResponseWriter, r *http.Request){
	
	// post data process
	r.ParseForm()
	email := r.FormValue("email")
	pswd := r.FormValue("pswd")

	//db open & close when whole function is executed
	db := connect()
	defer db.Close()

	//db process
	_, err := db.Query("INSERT INTO users(email, password) VALUES(?,?)", email, pswd)
	if err != nil {
		log.Print(err,email, pswd)
	}
}


//sign in handler
func signIn(w http.ResponseWriter, r *http.Request) {
	
	type M map[string]interface{} // storing tokenString
	
	//post data process
	r.ParseForm()
	email := r.FormValue("email")
	pswd := r.FormValue("pswd")

	//db open & close when whole function is executed
	db := connect()
	defer db.Close()

	//user auth
	row, err := db.Query("SELECT password FROM users WHERE email= ?", email)
	var userPass string //variable for password
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	
	for row.Next(){//converting db row into string data
		err := row.Scan(&userPass)
		if err != nil{
			log.Print(err)
		}
	}

	//generate token if user is logged in
	if pswd == userPass {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": email,
			"password": pswd,
		})
		signedToken, err := token.SignedString([]byte("sangatSecret"))
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		tokenString, _ := json.Marshal(M{"token" : signedToken})

		w.Write([]byte(tokenString)) //throw token to user
	}

}


func tokenAuth(next http.Handler) http.Handler{
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		

		if r.URL.Path == "/signin" || r.URL.Path =="/signup" { //check existing token on entry req 
            next.ServeHTTP(w, r)								// besides /signin and /signup
            return
        }
		//user's token auth
		authHeader := r.Header.Get("Authorization")
		
		if !strings.Contains(authHeader, "Bearer"){
			http.Error(w, "Invalid token", http.StatusBadRequest)
		}
	
		tokenString := strings.Replace(authHeader, "Bearer", "", -1) //slice for tokenstring only
	
		//parsing & validating user's token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Invalid signing method")
			}
			return []byte("sangatSecret"), nil
		})
		if err != nil && token == nil { //if this one is passed, token is verified
			http.Error(w, err.Error(), http.StatusBadRequest)
			return}
		 
		next.ServeHTTP(w, r)
	})
	

}






func tweet(w http.ResponseWriter, r *http.Request){
	r.ParseForm() //parse request data
	fmt.Println("alhamdulillah")
	fmt.Println(r.FormValue("tweet"))

	//db open & will close after user tweets
	db := connect()
	defer db.Close()

	_, err := db.Query("INSERT INTO tweets(email, time, tweet) VALUES(?, ?, ?)", r.FormValue("email"), time.Now(), r.FormValue("tweet"))
	if err != nil{
		log.Print(err)
	}
	w.Write([]byte("tweet posted"))
}