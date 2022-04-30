package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
	} else if command == "scene" {
		slug := os.Args[2]
		path := files.Filepath("scene.txt")
		os.Remove(path)
		cmd := exec.Command("vim", path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
		summary := files.ReadFile("scene.txt")
		fmt.Println(slug, summary)
	} else {
		fmt.Println("Command not available.")
	}

}
