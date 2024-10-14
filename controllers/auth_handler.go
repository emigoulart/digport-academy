package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// go get github.com/golang-jwt/jwt
// https://jwt.io/#debugger-io
var jwtKey = []byte("secret") // secret para a validação da assinatura

func validaToken(token string) (jwt.Claims, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return tkn.Claims, nil

}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Faltando authorization header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		_, err := validaToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Erro ao validar token"))
			return
		}

		//err = claims.Valid()
		//if err != nil {
		//	fmt.Fprintf(w, "Token not valid: %v", err)
		//}

		//email := claims.(jwt.MapClaims)["email"].(string)
		//role := claims.(jwt.MapClaims)["role"].(string)

		//r.Header.Set("email", email)
		//r.Header.Set("role", role)
		next.ServeHTTP(w, r)
	})
}

func GerarToken(w http.ResponseWriter) {
	dataExpiracao := time.Now().Add(60 * time.Minute)
	standardToken := &jwt.StandardClaims{
		ExpiresAt: dataExpiracao.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardToken)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println("Erro ao validar jwt:")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(tokenString))
}
