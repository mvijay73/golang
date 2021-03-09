package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var user Users
	json.NewDecoder(jsonFile).Decode(&user)

	fmt.Println(user)

}
