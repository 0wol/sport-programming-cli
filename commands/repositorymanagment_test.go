package commands

import (
	"fmt"
	"testing"
)

func TestFunctionsFromDirectory(t *testing.T) {
	res, _ := functionsFromDirectory("/Users/nikitaglusin/Documents/Coding/go/src/github.com/0w0l/sport-programming-cli/sport-programming-library")
	fmt.Println(res)
}
