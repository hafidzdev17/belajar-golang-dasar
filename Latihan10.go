package main

import "fmt"

// TODO: function type declaration
type Login func(params1, params2 string)string

func getLogin(username, password string, login Login){
	credentials := login(username,password)
	fmt.Println(credentials)
}

func setLogin(username, password string)string{
	if username == "hafid" && password == "123" {
		return "login success!"
	} else {
		if username == "" || password == "" {
			return "please input field username & password"
		} 
		if username == "sobri" && password == "123" {
			return "access denied!"
		}
		return "login failed!"
}
	}

func main(){
	getLogin("sobri","123",setLogin)
}