package file

import (
	"fmt"
	"testing"
)

func f(file *File, h int) {
	for i := 0; i < h; i++ {
		fmt.Print(" ")
	}
	fmt.Println(file.Name)
	for _, ff := range file.files {
		f(ff, h+4)
	}
}

func TestNewFile(t *testing.T) {
	absPath := "/Users/nikitaglusin/Documents/Coding/go/src/github.com/0w0l/sport-programming-cli/sport-programming-cli/sport-programming-library"

	file, err := NewFile("sport-programming-library", absPath)

	if err != nil {
		t.Error(err)
	}

	tree := file.getTree(0)

	// if err != nil {
	// 	t.Error(err)
	// }

	fmt.Println(tree)
}
