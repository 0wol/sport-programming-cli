package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	Repository     = "https://github.com/0wol/sport-programming-library.git"
	RepositoryName = "sport-programming-library"
	FunctionsFile  = "functions.fn"
)

func Update(ctx *cli.Context) error {
	err := updateRepository()

	if err != nil {
		return errors.New("An error was occurred")
	}

	return nil
}

func Get(ctx *cli.Context) error {
	if ctx.NArg() == 0 || ctx.NArg() > 3 {
		return errors.New("Invalid arguments count")
	}

	file, err := os.Open("./" + FunctionsFile)

	if err != nil {
		return errors.New("Please, update library by [update]")
	}
	defer file.Close()

	filter, err := getFilter(ctx)

	if err != nil {
		return err
	}

	function, err := getFunction(file, ctx.Args().First(), filter)

	if err != nil {
		return err
	}

	fmt.Println(function)
	return nil
}

func getFilter(ctx *cli.Context) (string, error) {
	var filter string
	if ctx.NArg() > 1 {
		if ctx.NArg() != 3 {
			return "", errors.New("Flag needs an argument [file, function]")
		}

		flag := ctx.Args().Get(1)
		if flag[0] != '-' {
			return "", errors.New("Invalid flag")
		}

		filter = ctx.Args().Get(2)

		if filter != "function" && filter != "file" {
			return "", errors.New("Invalid argument. Use [file, function]")
		}
	}

	return filter, nil
}
