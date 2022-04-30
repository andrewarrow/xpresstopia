package commands

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"xpresstopia/files"
)

type Group struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	From    int    `json:"from"`
	To      int    `json:"to"`
}

func AddGroup(name, from, to string) {
	path := files.Filepath("group.txt")
	os.Remove(path)
	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	summary := files.ReadFile("group.txt")

	play := files.ReadFile("screenplay.txt")

	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)

	list := []Group{}
	group := Group{name, strings.TrimSpace(summary), fromInt, toInt}
	list = append(list, group)
	b, _ := json.Marshal(list)
	files.SaveFile(play+"_groups.txt", string(b))
}