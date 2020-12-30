package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type user struct {
	Name     string
	Password string
}

func main() {

	login := flag.NewFlagSet("login", flag.ExitOnError)

	flag.Usage = func() {

		documentacion := `Las opciones son
login Para ingresar`

		fmt.Fprintf(os.Stderr, "%s\n", documentacion)
	}

	if len(os.Args) == 1 {
		flag.Usage()
	}

	for i, v := range os.Args {

		fmt.Printf("%v. %v\n", i, v)

	}

	switch os.Args[1] {
	case "login":
		username := login.String("username", "", "Introduce el username que desees")
		password := login.String("password", "", "Introduce la password que desees")
		login.Parse(os.Args[2:])
		fmt.Println(*username, *password)
	}

	archivo, err := os.OpenFile("databases.json", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

}
