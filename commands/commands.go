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
	GreenColor     = "\033[0;32m"
	RedColor       = "\033[0;31m"
	YellowColor    = "\033[1;33m"
)

func Update(ctx *cli.Context) error {
	var (
		err             error
		printableResult string
	)

	if !repositoryExists() {
		fmt.Printf("%sDownloading...\n", YellowColor)
		err = cloneRepository()
		printableResult = "downloaded"
	} else {
		fmt.Printf("%sUpdating...\n", YellowColor)
		err = pullRepository()
		printableResult = "updated"
	}

	if err != nil {
		return errors.New(fmt.Sprintf("%sAn error was occurred", RedColor))
	}

	fmt.Printf("%sLibrary was successfully %s\n", GreenColor, printableResult)

	return nil
}

func Get(ctx *cli.Context) error {
	if ctx.NArg() == 0 || ctx.NArg() > 3 {
		return errors.New(fmt.Sprintf("%sInvalid arguments count", RedColor))
	}

	if !repositoryExists() {
		return errors.New(fmt.Sprintf("%sPlease, update library by [update]", RedColor))
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
			return "", errors.New(fmt.Sprintf("%sFlag needs an argument [file, function]", RedColor))
		}

		flag := ctx.Args().Get(1)
		if flag[0] != '-' {
			return "", errors.New(fmt.Sprintf("%sInvalid flag", RedColor))
		}

		filter = ctx.Args().Get(2)

		if filter != "function" && filter != "file" {
			return "", errors.New(fmt.Sprintf("%sInvalid argument. Use [file, function]", RedColor))
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
