package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Welcome(w http.ResponseWriter,r *http.Request){
	c := r.Cookies()
	value := ""
	for _,cookie := range c{
		if cookie.Name == "token"{
			value = cookie.Value
		}
	}
	if value == ""{
		fmt.Println("err: missing token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	tkn,err := jwt.ParseWithClaims(value,claims,func(token *jwt.Token)(interface{},error){
		return jwtKey,nil
	})
	if err!=nil{
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("err:",err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("err:",err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		fmt.Println("err:token 不合法")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!",claims.Username)))


}