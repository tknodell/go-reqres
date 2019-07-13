package main

import (
	"fmt"

	"gopkg.in/resty.v1"
)

type Users struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
}

type User struct {
	Data struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
}

func main() {
	users := getAllUsers()

	// Print all users first & last name
	for _, user := range users.Data {
		fmt.Printf("%v %v\n", user.FirstName, user.LastName)
	}

	someUser := getUser("2")
	fmt.Println(someUser.Data.FirstName)
}

func getAllUsers() *Users {
	url := "https://reqres.in/api/users"
	var u Users
	resty.R().
		SetResult(&u).
		Get(url)

	return &u
}

func getUser(ID string) *User {
	url := fmt.Sprintf("https://reqres.in/api/users/%v", ID)
	var u User
	resty.R().
		SetResult(&u).
		Get(url)

	return &u
}
