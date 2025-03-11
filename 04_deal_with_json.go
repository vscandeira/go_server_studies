package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type socialMedia struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
type User struct {
	Id     int         `json:"id"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Age    int         `json:"age"`
	Social socialMedia `json:"social"`
}

func main() {
	jsonFile, err := os.Open("./04_sample_users.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened 04_sample_users.json")
	fmt.Println()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users []User
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users); i++ {
		fmt.Println("User id: ", users[i].Id)
		fmt.Println("User Name: ", users[i].Name)
		fmt.Println("User Email: ", users[i].Email)
		fmt.Println("User Age: ", users[i].Age)
		fmt.Println("User Social Media")
		fmt.Println("Facebook: ", users[i].Social.Facebook)
		fmt.Println("Twitter: ", users[i].Social.Twitter)
		fmt.Println()
	}
}
