package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"xpresstopia/files"
)

type Scene struct {
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
}

func ListScenes(play string) {
	existing := files.ReadFile(play + "_scenes.txt")
	list := []Scene{}
	if existing != "" {
		json.Unmarshal([]byte(existing), &list)
	}
	fmt.Println("")
	for i, s := range list {
		fmt.Printf("%2d. %s\n", i+1, s.Slug)
	}
	fmt.Println("")
}

func AddScene(play, slug, summary string) {
	existing := files.ReadFile(play + "_scenes.txt")
	list := []Scene{}
	if existing != "" {
		json.Unmarshal([]byte(existing), &list)
	}
	scene := Scene{slug, strings.TrimSpace(summary)}
	list = append(list, scene)
	b, _ := json.Marshal(list)
	files.SaveFile(play+"_scenes.txt", string(b))
}
