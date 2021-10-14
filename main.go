package main

import (
	"TEST_JWT/router"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/signin",router.Signin)
	http.HandleFunc("/welcome",router.Welcome)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080",nil))
}
