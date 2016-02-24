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
			log.Fatal(err)
		}
		fmt.Println(username, access_key)
	case "show":
		fmt.Println("show")
	case "remove":
		fmt.Println("delete")
	case "suspend":
		fmt.Println("suspend")
	}
}
