package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func SaveFile(name, data string) {
	home := UserHomeDir()
	fname := fmt.Sprintf("%s/%s/%s", home, STORAGE, name)
	os.Remove(fname)
	ioutil.WriteFile(fname, []byte(data), 0644)
}
