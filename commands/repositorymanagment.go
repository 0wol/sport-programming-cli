package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func updateRepository() error {
	exists, err := repositoryExists()

	if err != nil {
		return err
	}

	if !exists {
		err = cloneRepository()
	} else {
		err = pullRepository()
	}

	if err != nil {
		return err
	}

	err = saveFunctionsListToFile()

	if err != nil {
		return err
	}

	fmt.Println("Library was successfully updated")

	return nil
}

func cloneRepository() error {
	cmd := exec.Command("git", "clone", Repository)
	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	if strings.Contains(string(stdout), "not") {
		return errors.New("git is not installed")
	}

	return nil
}

func pullRepository() error {
	cmd := exec.Command("git", "-C", "./"+RepositoryName, "pull")
	_, err := cmd.Output()

	return err
}

func repositoryExists() (bool, error) {
	files, err := ioutil.ReadDir("./")

	if err != nil {
		return false, err
	}

	for _, f := range files {
		if f.Name() == RepositoryName {
			return true, nil
		}
	}

	return false, err
}

func saveFunctionsListToFile() error {
	path, err := filepath.Abs("./" + RepositoryName)

	if err != nil {
		return err
	}

	res, err := functionsFromDirectory(path)

	if err != nil {
		return err
	}

	file, err := os.Create(FunctionsFile)

	if err != nil {
		return err
	}

	file.WriteString(res)

	return nil
}

func functionsFromDirectory(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return "", err
	}

	res := ""

	for _, f := range files {
		if f.IsDir() {
			if f.Name() == ".git" {
				continue
			}
			functions, err := functionsFromDirectory(dir + "/" + f.Name())
			if err == nil {
				res += functions
			}
		} else {
			fn, err := functionsFromFile(dir + "/" + f.Name())
			if err == nil {
				res += fn
			}
		}
	}
	return res, nil
}

func functionsFromFile(filepath string) (string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	return parseFile(file), nil
}

func parseFile(file *os.File) string {
	scanner := bufio.NewScanner(file)

	res := ""
	level := 0
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(line, "{") {
			if level == 0 {
				fn := strings.Trim(line, "{ \t")
				res += fmt.Sprintf("%s~%s~%d\n", fn, file.Name(), lineNumber)
			}
			level++
		} else if strings.Contains(line, "}") {
			level--
		}
	}
	return res
}
