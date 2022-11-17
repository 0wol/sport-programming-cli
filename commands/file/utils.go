package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isFile(absPath string) bool {
	_, err := ioutil.ReadDir(absPath)
	if err != nil {
		return true
	}
	return false
}

func GetFileNameByAbsPath(absPath string) string {
	f, err := os.Open(absPath)
	if err != nil {
		return ""
	}
	return f.Name()
}

func getFiles(dir string) ([]*File, error) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var res []*File

	for _, f := range files {
		if f.IsDir() {
			if f.Name() == ".git" {
				continue
			}

			files, err := getFiles(dir + "/" + f.Name())

			if err != nil {
				return nil, err
			}

			res = append(res, &File{
				Name:         f.Name(),
				IsDirectory:  true,
				AbsolutePath: dir + "/" + f.Name(),
				files:        files,
			})
		} else {
			if !strings.Contains(f.Name(), ".cpp") {
				continue
			}
			res = append(res, &File{
				Name:         f.Name(),
				IsDirectory:  false,
				AbsolutePath: dir + "/" + f.Name(),
				files:        []*File{},
			})
		}
	}
	return res, nil
}

func parseFile(file *os.File) []string {
	var (
		scanner    = bufio.NewScanner(file)
		level      = 0
		lineNumber = 0
		res        []string
	)

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(line, "{") {
			if level == 0 {
				fn := strings.Trim(line, "{ \t")
				res = append(res, fmt.Sprintf("%s~%s~%d\n", getSignatureName(fn), file.Name(), lineNumber))
			}
			level++
		} else if strings.Contains(line, "}") {
			level--
		}
	}
	return res
}

func getSignatureName(line string) string {
	t := strings.Split(line, " ")
	if len(t) == 2 { // it is struct
		return t[1]
	}

	i := 0
	for line[i] != '(' {
		i++
	}

	i--
	for line[i] == ' ' {
		i--
	}

	res := ""

	for line[i] != ' ' {
		res = string(line[i]) + res
		i--
	}

	return res
}

func SliceContainsSimilar(slice []string, element string) bool {
	for _, e := range slice {
		if strings.Contains(element, e) {
			return true
		}
	}
	return false
}
