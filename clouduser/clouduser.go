package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Println("usage: clouduser <command> [<args>]")
	fmt.Println("examples:")
	fmt.Println("clouduser make janedoe")
	fmt.Println("clouduser show janedoe")
	fmt.Println("clouduser suspend janedoe")
	fmt.Println("clouduser enable janedoe")
	fmt.Println("clouduser delete janedoe")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	username := os.Args[2]

	switch os.Args[1] {
	case "make":
		access_key, err := insertUser(username)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(username, access_key)
		return
	case "show":
		user, err := getUser(username)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(user)
		return
	case "delete":
		err := deleteUser(username)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("User deleted.")
		return
	case "suspend":
		fmt.Println("suspend")
		return
	case "enable":
		fmt.Println("enable")
		return
	}
	usage()
	return
}
