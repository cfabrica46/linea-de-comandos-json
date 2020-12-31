package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type user struct {
	Name     string
	Password string
}

func main() {

	var data []byte
	var users []user
	var acceso bool

	login := flag.NewFlagSet("login", flag.ExitOnError)
	register := flag.NewFlagSet("register", flag.ExitOnError)

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

	archivo, err := os.OpenFile("databases.json", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	buf := make([]byte, 512)

	for {
		n, err := archivo.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		data = append(data, buf[:n]...)
	}

	err = json.Unmarshal(data, &users)

	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "login":
		username := login.String("username", "", "Introduce el username que desees")
		password := login.String("password", "", "Introduce la password que desees")
		login.Parse(os.Args[2:])

		for i := range users {

			if users[i].Name == *username && users[i].Password == *password {

				acceso = true

			}

		}

		if acceso == true {
			fmt.Println("Bienvenido", *username)
		} else {
			fmt.Println("Username y/o Password incorrectas")
		}

	case "register":
		username := register.String("username", "", "Introduce el username que desees")
		password := register.String("password", "", "Introduce la password que desees")
		register.Parse(os.Args[2:])

		nuevoUser := user{Name: *username, Password: *password}

		users = append(users, nuevoUser)

		nuevaData, err := json.MarshalIndent(users, "", " ")

		if err != nil {
			log.Fatal(err)
		}

		archivo.Seek(0, 0)

		_, err = archivo.Write(nuevaData)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Nuevo usuario registrado", username)
	}

}
