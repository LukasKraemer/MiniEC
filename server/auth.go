package main

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"strings"
	"time"
)
type User struct {
	IdUser uint64   `json:"id"`
	Email string 	`json:"username"`
	Permission int	 `json:"permission"`
	Password string `json:"password"`
}


func login(w http.ResponseWriter, r *http.Request) {
	stat, valid := CreateToken(r)
	if !valid{
		http.Error(w, "authorization failed", http.StatusUnauthorized)
	}else{
		json.NewEncoder(w).Encode(stat)
	}
}

func validate(username string, password string) (uint64, int) {
	user := User{}

	h := sha512.New()
	h.Write([]byte(password))
	hash := hex.EncodeToString(h.Sum(nil))

	db, _ := sql.Open("mysql", os.Getenv("databaseString"))
	_ = db.QueryRow ("SELECT idUser, email, password, permission FROM user where password = ? and email = ? and permission > 0;", hash, username).Scan(&user.IdUser, &user.Email, &user.Password, &user.Permission)
	defer db.Close()
	return user.IdUser, user.Permission
}



func CreateToken(r *http.Request) (string, bool) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return "", false
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) == 2{
		userid, permission := validate(pair[0], pair[1])
		if permission != -1 && userid != 0 {
			os.Getenv("ACCESS_SECRET") //this should be in an env file
			atClaims := jwt.MapClaims{}
			atClaims["permission"] = permission
			atClaims["user_id"] = userid
			atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
			if err == nil {
				return token, true
			}
		}
	}
	return "", false
}