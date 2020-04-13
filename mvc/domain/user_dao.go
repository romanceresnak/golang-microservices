package domain

import (
	"github.com/romanceresnak/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{123 : &User{
		Id:        123,
		FirstName: "Peter",
		LastName:  "Klacansky",
		Email:     "peter.klacansky@gmail.com",
	},
	}
)

//if i use pointer user i do not have to define new variable
func GetUser(userId int64) (*User,*utils.ApplicationError){
	if user := users[userId]; user != nil{
		return user,nil
	}
	return nil, &utils.ApplicationError{
		Message: "User was not found",
		StatusCode: http.StatusNotFound ,
		Code:    "not_found",
	}
}
