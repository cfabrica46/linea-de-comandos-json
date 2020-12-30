package main

import (
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

}
