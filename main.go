package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed ace.jpg
var ace []byte

//go:embed file.txt
var version string

//go:embed files/*.txt
var Multiple embed.FS

func main() {
	err := ioutil.WriteFile("ace_clone.jpg", ace, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(version)
	dir, err := Multiple.ReadDir("files")
	if err != nil {
		panic(err)
	}
	for _, list := range dir {
		if !list.IsDir() {
			fmt.Println(list.Name())
		}
	}
}
