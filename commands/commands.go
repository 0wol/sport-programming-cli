package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"sport-programming-cli/commands/file"

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

	if !repositoryExists() {
		return errors.New("Please, update library by [update]")
	}

	filter, err := getFilter(ctx)

	if err != nil {
		return err
	}

	absPath, err := filepath.Abs("./" + RepositoryName)

	if err != nil {
		return err
	}

	dir, err := file.NewFile(RepositoryName, absPath)

	if err != nil {
		return err
	}

	var (
		bodies string
		query  = strings.ToLower(ctx.Args().First())
	)

	if filter == "file" {
		bodies, err = dir.FindFunctionsByFileName(query)

		if err != nil {
			return err
		}
	} else {
		bodies, err = dir.FindFunctionsByName(query)

		if err != nil {
			return err
		}
	}

	fmt.Println(strings.TrimSuffix(bodies, "\n\n"))

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

func List(ctx *cli.Context) error {
	absPath, err := filepath.Abs("./" + RepositoryName)

	if err != nil {
		return err
	}

	dir, err := file.NewFile(RepositoryName, absPath)

	if err != nil {
		return err
	}

	fmt.Println(strings.TrimSuffix(dir.GetTree(), "\n"))

	return nil
}
