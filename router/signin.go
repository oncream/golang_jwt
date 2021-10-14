package router

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("FSMTjRoFVOzugwDJgxTPlVIEEqUYqyhJ")

var users = map[string]string{
	"user1":"password1",
	"user2":"password2",
}
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter,r *http.Request){
	fmt.Println("signin")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds);err!=nil{
		fmt.Println("err:",err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword,ok := users[creds.Username]
	if !ok || expectedPassword!=creds.Password{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5*time.Minute)
	claims := &Claims{
		Username:creds.Username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwt.New(jwt.SigningMethodHS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil{
		fmt.Println("err:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,&http.Cookie{
		Name:"token",
		Value:tokenString,
		Expires: expirationTime,
	})
}