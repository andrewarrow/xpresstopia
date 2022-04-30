package files

import (
	"fmt"
	"os"
	"runtime"
)

const STORAGE = ".expresstopia"

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func MakeStorageDir() {
	home := UserHomeDir()
	os.Mkdir(fmt.Sprintf("%s/%s", home, STORAGE), 0755)
}

func Filepath(name string) string {
	home := UserHomeDir()
	fname := fmt.Sprintf("%s/%s/%s", home, STORAGE, name)
	return fname
}
