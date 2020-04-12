package finder

import (
	"io/ioutil"
)

type MyFile struct {
	FileName string
}

func (f MyFile) GetContent() (string, error) {
	content, err := ioutil.ReadFile(f.FileName)
	if err != nil {
		return "", err
	}
	return string(content), nil;
}