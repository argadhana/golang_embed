package test

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed ace.jpg
var ace []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("ace_clone.jpg", ace, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
