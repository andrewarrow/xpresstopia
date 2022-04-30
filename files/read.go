package files

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(name string) string {
	home := UserHomeDir()
	fname := fmt.Sprintf("%s/%s/%s", home, STORAGE, name)
	b, _ := ioutil.ReadFile(fname)
	return string(b)
}
