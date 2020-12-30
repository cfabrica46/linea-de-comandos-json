package main

import (
	"encoding/json"
	"log"
	"os"
)

type user struct {
	Name     string
	Password string
}

func main() {

	archivo, err := os.OpenFile("databases.json", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	user1 := user{Name: "cfabrica46", Password: "cesaruwu"}

	user2 := user{Name: "carlos", Password: "carlosxd"}

	users := []user{user1, user2}

	data, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	archivo.Write(data)
}
