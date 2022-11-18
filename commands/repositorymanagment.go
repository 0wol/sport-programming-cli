package commands

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

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

func repositoryExists() bool {
	files, err := ioutil.ReadDir("./")

	if err != nil {
		return false
	}

	for _, f := range files {
		if f.Name() == RepositoryName {
			return true
		}
	}

	return false
}
