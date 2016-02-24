package main

import (
	"fmt"
	"log"
	"os"
)

func Usage() {
	fmt.Println("usage: clouduser <command> [<args>]")
	fmt.Println("examples:")
	fmt.Println("clouduser make janedoe")
	fmt.Println("clouduser show janedoe")
	fmt.Println("clouduser suspend janedoe")
	fmt.Println("clouduser enable janedoe")
	fmt.Println("clouduser delete janedoe")
}

func Make(username string) {
	access_key, err := insertUser(username)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(username, access_key)
}

func Show(username string) {
	user, err := getUser(username)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(user)
}

func Delete(username string) {
	err := deleteUser(username)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("User deleted.")
}

func Suspend(username string) {
	fmt.Println("suspend")
}

func Enable(username string) {
	fmt.Println("enable")
}

func main() {
	if len(os.Args) < 3 {
		Usage()
		os.Exit(1)
	}
	username := os.Args[2]

	switch os.Args[1] {
	case "make":
		Make(username)
		os.Exit(0)
	case "show":
		Show(username)
		os.Exit(0)
	case "delete":
		Delete(username)
		os.Exit(0)
	case "suspend":
		Suspend(username)
		os.Exit(0)
	case "enable":
		Enable(username)
		os.Exit(0)
	}
	Usage()
	os.Exit(1)
}
