package test

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var Multiple embed.FS

func TestMultipleEmbed(t *testing.T) {
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
