package commands

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getFunction(file *os.File, function string, filter string) (string, error) {
	scanner := bufio.NewScanner(file)

	res := ""

	for scanner.Scan() {
		line := scanner.Text()
		t := strings.Split(line, "~")

		if (strings.Contains(t[0], function) && (filter == "" || filter == "function")) || (strings.Contains(getRepositoryDirectory(t[1]), function) && filter == "file") {
			funcFile, err := os.Open(t[1])
			if err != nil {
				continue
			}
			lineNumber, err := strconv.Atoi(t[2])
			if err != nil {
				continue
			}
			body := getFunctionBody(funcFile, lineNumber)
			funcFile.Close()

			res += body + "\n\n"
		}
	}

	res = strings.TrimSuffix(res, "\n\n")

	return res, nil
}

func getRepositoryDirectory(absPath string) string {
	var (
		t = strings.Split(absPath, "/")
		i int
	)

	for i = 0; i < len(t) && t[i] != RepositoryName; i++ {
	}

	return strings.Join(t[i:], "/")
}

func getFunctionBody(file *os.File, lineNumber int) string {
	var (
		scanner           = bufio.NewScanner(file)
		res               = ""
		currentLineNumber = 0
		level             = 0
		prev              = ""
	)

	for scanner.Scan() {
		currentLineNumber++
		line := scanner.Text()
		if currentLineNumber == lineNumber {
			if prev != "" {
				res += prev + "\n"
			}
			res += line + "\n"
			level++
			for level > 0 && scanner.Scan() {
				line := scanner.Text()

				if strings.Contains(line, "{") {
					level++
				} else if strings.Contains(line, "}") {
					level--
				}
				res += line + "\n"
			}
		}
		prev = line
	}

	res = strings.TrimSuffix(res, "\n")

	return res
}
