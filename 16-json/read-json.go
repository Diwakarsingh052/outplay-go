package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Name  string          `json:"first_name"`
	Perms map[string]bool `json:"perms"`
}

func main() {

	jsonData, err := os.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}
	var users []user

	err = json.Unmarshal(jsonData, &users) // Unmarshal needs reference of your data

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(users)
	fmt.Printf("\n%#v", users)

}
