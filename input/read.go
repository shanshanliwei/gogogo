package input

import (
	"io/ioutil"
)

func GetFileStringByPath(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}
