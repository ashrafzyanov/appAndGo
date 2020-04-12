package main

import (
	"fmt"
	"github.com/ashrafzyanov/appAndGo/finder"
)

func main() {
	filePath := getFilePath();
	var contentGetter finder.ContentGetter;
	contentGetter = finder.MyFile{FileName: filePath}
	content, err := contentGetter.GetContent();
	if err != nil {
		fmt.Println("SomeProblem: ", err.Error())
		return
	}
	fmt.Println("File Content: \n" + content);
}

func getFilePath() string {
	var file string;
	fmt.Scanf("%s", &file)
	return file;
}