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
		return
	}
	fmt.Println(username, access_key)
}

func Show(username string) {
	user, err := getUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(user)
}

func Delete(username string) {
	err := deleteUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("User deleted.")
}

func Suspend(username string) {
	err := setActive(username, 0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("User suspended.")
	}
}

func Enable(username string) {
	err := setActive(username, 1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("User enabled.")
	}
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
	case "show":
		Show(username)
	case "delete":
		Delete(username)
	case "suspend":
		Suspend(username)
	case "enable":
		Enable(username)
	default:
		Usage()
	}
}
