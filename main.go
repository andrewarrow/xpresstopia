package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"xpresstopia/commands"
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
		files.SaveFile("screenplay.txt", name)
	} else if command == "ls" {
		play := files.ReadFile("screenplay.txt")
		flavor := "scenes"
		if len(os.Args) > 2 {
			flavor = os.Args[2]
		}
		if flavor == "scenes" {
			commands.ListScenes(play)
		} else {
			commands.ListGroups(play)
		}
	} else if command == "group" {
		group := os.Args[2]
		from := os.Args[3]
		to := os.Args[4]
		commands.AddGroup(group, from, to)
	} else if command == "scene" {
		slug := os.Args[2]
		path := files.Filepath("scene.txt")
		os.Remove(path)
		cmd := exec.Command("vim", path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
		play := files.ReadFile("screenplay.txt")
		summary := files.ReadFile("scene.txt")
		commands.AddScene(play, slug, summary)
	} else {
		fmt.Println("Command not available.")
	}

}
