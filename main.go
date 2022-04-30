package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"xpresstopia/files"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	files.MakeStorageDir()

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	//argMap := ArgsToMap()

	if command == "new" {
		name := os.Args[2]
		fmt.Println(name)
	} else if command == "foo" {
	} else {
		fmt.Println("Command not available.")
	}

}
