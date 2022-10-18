package key

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type RequestModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var jwtKey = []byte("secret_key")

func GetToken(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokenString)
}

// func GenerateJWT(email string, username string) (tokenString string, err error) {
// 	expirationTime := time.Now().Add(1 * time.Hour)
// 	claims := &JWTClaim{
// 		Email:    email,
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err = token.SignedString(jwtKey)
// 	fmt.Print(tokenString)
// 	return
// }

func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body) // get request body
	var reqModel RequestModel
	json.Unmarshal(reqBody, &reqModel)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    reqModel.Email,
		Username: reqModel.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tokenString)
}
