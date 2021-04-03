package test

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed file.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}
