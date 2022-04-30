package commands

import (
	"encoding/json"
	"strings"
	"xpresstopia/files"
)

type Scene struct {
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
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
